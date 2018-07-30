package mysqlx

import (
	"context"
	"crypto/tls"
	"database/sql"
	"database/sql/driver"
	"encoding/binary"
	"io"
	"net"
	"time"

	"github.com/renthraysk/mysqlx/authentication"
	"github.com/renthraysk/mysqlx/authentication/plain"
	"github.com/renthraysk/mysqlx/msg"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_notice"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_session"
	"github.com/renthraysk/mysqlx/slice"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type conn struct {
	netConn   net.Conn
	connector *Connector
	buf       []byte

	hasClientID bool
	clientID    uint64
}

func (c *conn) replaceBuffer() {
	n := cap(c.buf)
	if n < minBufferSize {
		n = minBufferSize
	}
	c.buf = make([]byte, n)
}

func (c *conn) readMessage(ctx context.Context) (mysqlx.ServerMessages_Type, []byte, error) {
	s := c.buf[:5]
	if _, err := io.ReadFull(c.netConn, s); err != nil {
		return 0, nil, err
	}
	n := binary.LittleEndian.Uint32(s)
	t := mysqlx.ServerMessages_Type(s[4])
	if n <= 1 {
		return t, nil, nil
	}
	_, s = slice.Allocate(c.buf[:0], int(n)-1)
	if _, err := io.ReadFull(c.netConn, s); err != nil {
		return 0, nil, err
	}
	return t, s, nil
}

func (c *conn) send(ctx context.Context, m msg.Msg) error {
	deadline, _ := ctx.Deadline()
	if err := c.netConn.SetDeadline(deadline); err != nil {
		return errors.Wrap(err, "unable to set deadline")
	}
	_, err := m.WriteTo(c.netConn)
	return err
}

func (c *conn) execMsg(ctx context.Context, m msg.Msg) (driver.Result, error) {
	if err := c.send(ctx, m); err != nil {
		return nil, err
	}
	r := &result{}
	for {
		t, b, err := c.readMessage(ctx)
		if err != nil {
			return nil, err
		}
		switch t {
		case mysqlx.ServerMessages_OK:
			return r, nil

		case mysqlx.ServerMessages_ERROR:
			return nil, newError(b)

		case mysqlx.ServerMessages_SQL_STMT_EXECUTE_OK:
			return r, nil

		case mysqlx.ServerMessages_NOTICE:
			var f mysqlx_notice.Frame

			if err := proto.Unmarshal(b, &f); err != nil {
				return nil, errors.Wrap(err, "failed to unmarshal Frame")
			}

			switch mysqlx_notice.Frame_Type(f.GetType()) {
			case mysqlx_notice.Frame_WARNING:
				var w mysqlx_notice.Warning

				if err := proto.Unmarshal(f.Payload, &w); err != nil {
					return nil, errors.Wrap(err, "failed to unmarshal Warning")
				}

			case mysqlx_notice.Frame_SESSION_VARIABLE_CHANGED:
				var v mysqlx_notice.SessionVariableChanged

				if err := proto.Unmarshal(f.Payload, &v); err != nil {
					return nil, errors.Wrap(err, "failed to unmarshal SessionVariableChanged")
				}

			case mysqlx_notice.Frame_SESSION_STATE_CHANGED:
				var s mysqlx_notice.SessionStateChanged

				if err := proto.Unmarshal(f.Payload, &s); err != nil {
					return nil, errors.Wrap(err, "failed to unmarshal SessionStateChanged")
				}
				switch s.GetParam() {
				case mysqlx_notice.SessionStateChanged_CURRENT_SCHEMA:
				case mysqlx_notice.SessionStateChanged_ACCOUNT_EXPIRED:
				case mysqlx_notice.SessionStateChanged_GENERATED_INSERT_ID:
					r.lastInsertID, r.hasLastInsertID = ScalarUint(s.Value)

				case mysqlx_notice.SessionStateChanged_ROWS_AFFECTED:
					r.rowsAffected, r.hasRowsAffected = ScalarUint(s.Value)

				case mysqlx_notice.SessionStateChanged_ROWS_FOUND:
					r.rowsFound, r.hasRowsFound = ScalarUint(s.Value)

				case mysqlx_notice.SessionStateChanged_ROWS_MATCHED:
					r.rowsMatched, r.hasRowsMatched = ScalarUint(s.Value)

				case mysqlx_notice.SessionStateChanged_TRX_COMMITTED:
				case mysqlx_notice.SessionStateChanged_TRX_ROLLEDBACK:
				case mysqlx_notice.SessionStateChanged_PRODUCED_MESSAGE:

				case mysqlx_notice.SessionStateChanged_CLIENT_ID_ASSIGNED:
					c.clientID, c.hasClientID = ScalarUint(s.Value)
				}
			}
		default:
		}
	}
}

func (c *conn) ExecContext(ctx context.Context, stmt string, args []driver.NamedValue) (driver.Result, error) {
	s, err := msg.StmtNamedValues(c.buf[:0], stmt, args)
	if err != nil {
		return nil, err
	}
	return c.execMsg(ctx, s)
}

func (c *conn) queryMsg(ctx context.Context, msg msg.Msg) (driver.Rows, error) {
	if err := c.send(ctx, msg); err != nil {
		return nil, err
	}
	r := &rows{
		conn: c,
	}
	if err := r.readColumns(ctx); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *conn) QueryContext(ctx context.Context, stmt string, args []driver.NamedValue) (driver.Rows, error) {
	s, err := msg.StmtNamedValues(c.buf[:0], stmt, args)
	if err != nil {
		return nil, err
	}
	return c.queryMsg(ctx, s)
}

func (c *conn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	return c.connector.stmtPreparer(ctx, c, query)
}

// Prepare driver.Conn interface forces this deprecated implementation
func (c *conn) Prepare(query string) (driver.Stmt, error) {
	return c.connector.stmtPreparer(context.Background(), c, query)
}

func (c *conn) Close() error {
	return c.closeContext(context.Background())
}

func (c *conn) closeContext(ctx context.Context) error {
	if c.netConn == nil {
		return nil
	}
	err := c.send(ctx, msg.ConnectionClose(c.buf[:0]))
	_, _, _ = c.readMessage(ctx)
	errClose := c.netConn.Close()
	c.netConn = nil
	if err != nil {
		return err
	}
	return errClose
}

func (c *conn) BeginTx(ctx context.Context, options driver.TxOptions) (driver.Tx, error) {
	set := ""
	start := "START TRANSACTION"
	if options.ReadOnly {
		start = "START TRANSACTION READ ONLY"
	}
	switch sql.IsolationLevel(options.Isolation) {
	case sql.LevelDefault:
	case sql.LevelReadUncommitted:
		set = "SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED"
	case sql.LevelReadCommitted:
		set = "SET TRANSACTION ISOLATION LEVEL READ COMMITTED"
	case sql.LevelRepeatableRead:
		set = "SET TRANSACTION ISOLATION LEVEL REPEATABLE READ"
	case sql.LevelSerializable:
		set = "SET TRANSACTION ISOLATION LEVEL SERIALIZABLE"
	case sql.LevelSnapshot:
		start = "START TRANSACTION WITH CONSISTENT SNAPSHOT"
		if options.ReadOnly {
			start = "START TRANSACTION WITH CONSISTENT SNAPSHOT, READ ONLY"
		}
	default:
		return nil, errors.Errorf("Unsupported transaction isolation level (%s)", sql.IsolationLevel(options.Isolation).String())
	}

	if len(set) > 0 {
		if _, err := c.ExecContext(ctx, set, nil); err != nil {
			return nil, err
		}
	}
	if _, err := c.ExecContext(ctx, start, nil); err != nil {
		return nil, err
	}
	return &tx{c}, nil
}

// Begin driver.Conn interface forces this deprecated implementation
func (c *conn) Begin() (driver.Tx, error) {
	if _, err := c.ExecContext(context.Background(), "START TRANSACTION", nil); err != nil {
		return nil, err
	}
	return &tx{c}, nil
}

func (c *conn) Ping(ctx context.Context) error {
	_, err := c.execMsg(ctx, msg.Ping(c.buf[:0]))
	return err
}

func (c *conn) ResetSession(ctx context.Context) error {
	return c.connector.sessionResetter(ctx, c)
}

func (c *conn) CheckNamedValue(nv *driver.NamedValue) error {
	if nv.Value == nil {
		return nil
	}
	switch v := nv.Value.(type) {
	case uint64, int64, string, []byte, float32, float64, bool, time.Time:
		// Protocol supported types.
	case uint8, uint16, uint32, uint, int8, int16, int32, int:
		// Supported via conversion to a type in above case.
	case time.Duration:
		const max = 838*time.Hour + 59*time.Minute + 59*time.Second
		if v > max {
			return errors.Errorf("time.Duration overflows mysql TIME (838:59:59)")
		}
		if v < -max {
			return errors.Errorf("time.Duration underflows mysql TIME (-838:59:59)")
		}
	default:
		if _, ok := nv.Value.(msg.ArgAppender); ok {
			return nil
		}
		return errors.Errorf("Unsupported type %T", nv.Value)
	}
	return nil
}

func (c *conn) enableTLS(ctx context.Context, tlsConfig *tls.Config) error {
	s := msg.NewCapabilitySetTLSEnable(c.buf[:0])
	if _, err := c.execMsg(ctx, s); err != nil {
		c.netConn.Close()
		return errors.Wrap(err, "failed to set TLS capability")
	}
	tlsConn := tls.Client(c.netConn, tlsConfig)
	if err := tlsConn.Handshake(); err != nil {
		tlsConn.Close()
		return errors.Wrap(err, "failed TLS handshake")
	}
	c.netConn = tlsConn
	return nil
}

func (c *conn) authenticate(ctx context.Context) error {

	const ER_ACCESS_DENIED_ERROR = 1045

	err := c.authenticate2(ctx, c.connector.authentication)
	if err == nil {
		return nil
	}
	if e, ok := errors.Cause(err).(*Error); !ok || e.Code != ER_ACCESS_DENIED_ERROR {
		return err
	}
	switch c.netConn.(type) {
	case *tls.Conn, *net.UnixConn:
		// Connected securely, so can attempt to authenticate with PLAIN,
		// which will populate the cache for caching_sha2 and sha256_password to start working
		if err2 := c.authenticate2(ctx, plain.New()); err2 == nil {
			return nil
		}
	default:
		// @TODO Need to decide what to do here..
		// https://dev.mysql.com/doc/refman/8.0/en/x-plugin-sha2-cache-plugin.html
		// Current feeling is to not allow authentication with sha2 over non secure connections,
		// as cannot initially populate the cache without TLS.
	}
	return err
}

func (c *conn) authenticate2(ctx context.Context, starter authentication.Starter) error {
	if err := c.send(ctx, starter.Start(c.buf[:0], c.connector)); err != nil {
		return err
	}
	for {
		t, b, err := c.readMessage(ctx)
		if err != nil {
			return errors.Wrap(err, "failed reading AuthenticateStart response")
		}
		switch t {
		case mysqlx.ServerMessages_NOTICE:
			continue

		case mysqlx.ServerMessages_ERROR:
			return newError(b)

		case mysqlx.ServerMessages_SESS_AUTHENTICATE_OK:
			return nil

		case mysqlx.ServerMessages_SESS_AUTHENTICATE_CONTINUE:
			continuer, ok := starter.(authentication.StartContinuer)
			if !ok {
				return errors.New("unexpected AuthenticateContinue")
			}

			var ac mysqlx_session.AuthenticateContinue
			if err := proto.Unmarshal(b, &ac); err != nil {
				return errors.Wrap(err, "failed to unmarshal AuthenticateContinue")
			}
			if err := c.send(ctx, continuer.Continue(c.buf[:0], c.connector, ac.AuthData)); err != nil {
				return errors.Wrap(err, "failed sending AuthenticateContinue")
			}

			for {
				t, b, err := c.readMessage(ctx)
				if err != nil {
					return errors.Wrap(err, "failed reading AuthenticateContinue response")
				}
				switch t {
				case mysqlx.ServerMessages_NOTICE:
					continue

				case mysqlx.ServerMessages_ERROR:
					return newError(b)

				case mysqlx.ServerMessages_SESS_AUTHENTICATE_OK:
					return nil

				default:
					return errors.Errorf("unexpected server response to AuthenticateContinue %s(%d)", t.String(), t)
				}
			}
		default:
			return errors.Errorf("unexpected server response to AuthenticateStart %s(%d)", t.String(), t)
		}
	}
}
