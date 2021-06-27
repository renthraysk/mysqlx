package mysqlx

import (
	"bufio"
	"context"
	"crypto/tls"
	"database/sql"
	"database/sql/driver"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"sync/atomic"
	"time"

	"github.com/renthraysk/mysqlx/authentication"
	"github.com/renthraysk/mysqlx/authentication/plain"
	"github.com/renthraysk/mysqlx/errs"
	"github.com/renthraysk/mysqlx/msg"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_notice"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_session"
	"google.golang.org/protobuf/proto"
)

type connStatus uint32

const (
	statusOK connStatus = iota
	statusBad
)

func (c *connStatus) Get() connStatus {
	return (connStatus)(atomic.LoadUint32((*uint32)(c)))
}

func (c *connStatus) Set(s connStatus) {
	atomic.StoreUint32((*uint32)(c), (uint32)(s))
}

type conn struct {
	status connStatus

	netConn   net.Conn
	r         *bufio.Reader
	discard   int
	connector *Connector
	buf       []byte

	hasClientID   bool
	clientID      uint64
	preparedStmts map[string]uint32
}

func (c *conn) replaceBuffer() {
	n := cap(c.buf)
	if n < minBufferSize {
		n = minBufferSize
	}
	c.buf = make([]byte, n)
}

func (c *conn) readMessage(ctx context.Context) (mysqlx.ServerMessages_Type, []byte, error) {
	if c.discard > 0 {
		c.r.Discard(c.discard)
		c.discard = 0
	}
	b, err := c.r.Peek(5)
	if err != nil {
		return 0, nil, err
	}
	n := binary.LittleEndian.Uint32(b)
	t := mysqlx.ServerMessages_Type(b[4])
	c.r.Discard(5)
	if n <= 1 {
		return t, nil, nil
	}
	n--
	switch b, err = c.r.Peek(int(n)); err {
	case bufio.ErrBufferFull:
		if cap(b) < int(n) {
			c.buf = make([]byte, (n+4095) & ^uint32(4095))
		}
		i := copy(c.buf, b)
		c.r.Discard(i)
		b = c.buf[:n]
		_, err = io.ReadFull(c.r, b[i:])
	case nil:
		c.discard = int(n)

	default:
	}
	return t, b, err
}

func (c *conn) send(ctx context.Context, m msg.Msg) error {
	deadline, _ := ctx.Deadline()
	if err := c.netConn.SetDeadline(deadline); err != nil {
		return fmt.Errorf("unable to set deadline: %w", err)
	}
	n, err := m.WriteTo(c.netConn)
	if err != nil && n == 0 {
		c.status.Set(statusBad)
	}
	return err
}

func (c *conn) sendN(ctx context.Context, b []byte) error {
	if len(b) == 0 {
		return nil
	}
	deadline, _ := ctx.Deadline()
	if err := c.netConn.SetDeadline(deadline); err != nil {
		return fmt.Errorf("unable to set deadline: %w", err)
	}
	if _, err := c.netConn.Write(b); err != nil {
		return fmt.Errorf("sending: %w", err)
	}
	var err errs.Errors

	// @TODO Needs to understand expectations, and mimick error handling
	for i := 0; len(b) >= 5; i++ {
	read:
		for {
			t, s, err2 := c.readMessage(ctx)
			if err2 != nil {
				return err2
			}
			switch t {
			case mysqlx.ServerMessages_OK, mysqlx.ServerMessages_SQL_STMT_EXECUTE_OK:
				break read

			case mysqlx.ServerMessages_NOTICE:

			case mysqlx.ServerMessages_ERROR:
				if err == nil {
					err = make(errs.Errors)
				}
				err[i] = errs.New(s)
				break read
			}
		}
		b = b[binary.LittleEndian.Uint32(b)+4:]
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (c *conn) execMsg(ctx context.Context, m msg.Msg) (driver.Result, error) {
	switch c.status.Get() {
	case statusBad:
		return nil, driver.ErrBadConn
	}
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
			return nil, c.handleError(b)

		case mysqlx.ServerMessages_SQL_STMT_EXECUTE_OK:
			return r, nil

		case mysqlx.ServerMessages_NOTICE:
			var f mysqlx_notice.Frame

			if err := proto.Unmarshal(b, &f); err != nil {
				return nil, fmt.Errorf("failed to unmarshal Frame: %w", err)
			}

			switch mysqlx_notice.Frame_Type(f.GetType()) {
			case mysqlx_notice.Frame_WARNING:
				var w mysqlx_notice.Warning

				if err := proto.Unmarshal(f.Payload, &w); err != nil {
					return nil, fmt.Errorf("failed to unmarshal Warning: %w", err)
				}

			case mysqlx_notice.Frame_SESSION_VARIABLE_CHANGED:
				var v mysqlx_notice.SessionVariableChanged

				if err := proto.Unmarshal(f.Payload, &v); err != nil {
					return nil, fmt.Errorf("failed to unmarshal SessionVariableChanged: %w", err)
				}

			case mysqlx_notice.Frame_SESSION_STATE_CHANGED:
				var s mysqlx_notice.SessionStateChanged

				if err := proto.Unmarshal(f.Payload, &s); err != nil {
					return nil, fmt.Errorf("failed to unmarshal SessionStateChanged: %w", err)
				}
				switch s.GetParam() {
				case mysqlx_notice.SessionStateChanged_CURRENT_SCHEMA:
				case mysqlx_notice.SessionStateChanged_ACCOUNT_EXPIRED:
				case mysqlx_notice.SessionStateChanged_GENERATED_INSERT_ID:

					r.lastInsertID, r.hasLastInsertID = scalarUint(s.Value[0])

				case mysqlx_notice.SessionStateChanged_ROWS_AFFECTED:
					if len(s.Value) != 1 {
						return nil, errors.New("unexpected number of rows affected values")
					}
					r.rowsAffected, r.hasRowsAffected = scalarUint(s.Value[0])

				case mysqlx_notice.SessionStateChanged_ROWS_FOUND:
					if len(s.Value) != 1 {
						return nil, errors.New("unexpected number of rows found values")
					}
					r.rowsFound, r.hasRowsFound = scalarUint(s.Value[0])

				case mysqlx_notice.SessionStateChanged_ROWS_MATCHED:
					if len(s.Value) != 1 {
						return nil, errors.New("unexpected number of rows matched values")
					}
					r.rowsMatched, r.hasRowsMatched = scalarUint(s.Value[0])

				case mysqlx_notice.SessionStateChanged_TRX_COMMITTED:
				case mysqlx_notice.SessionStateChanged_TRX_ROLLEDBACK:

				case mysqlx_notice.SessionStateChanged_PRODUCED_MESSAGE:

				case mysqlx_notice.SessionStateChanged_CLIENT_ID_ASSIGNED:
					c.clientID, c.hasClientID = scalarUint(s.Value[0])
				}
			}
		default:
		}
	}
}

func (c *conn) ExecContext(ctx context.Context, stmt string, args []driver.NamedValue) (driver.Result, error) {
	s, err := msg.NewStmtExecuteNamed(c.buf[:0], stmt, args)
	if err != nil {
		return nil, err
	}
	return c.execMsg(ctx, s)
}

func (c *conn) queryMsg(ctx context.Context, msg msg.Msg) (driver.Rows, error) {
	switch c.status.Get() {
	case statusBad:
		return nil, driver.ErrBadConn
	}
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
	s, err := msg.NewStmtExecuteNamed(c.buf[:0], stmt, args)
	if err != nil {
		return nil, err
	}
	return c.queryMsg(ctx, s)
}

func (c *conn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	if c.preparedStmts == nil {
		c.preparedStmts = make(map[string]uint32)
	}
	id, ok := c.preparedStmts[query]
	if !ok {
		id = uint32(len(c.preparedStmts) + 1)
		c.preparedStmts[query] = id
	}

	if _, err := c.execMsg(ctx, msg.NewPrepare(c.buf[:0], id, query)); err != nil {
		return nil, err
	}

	return &stmt{
		c:  c,
		id: id,
	}, nil
}

// Prepare driver.Conn interface forces this deprecated implementation
func (c *conn) Prepare(query string) (driver.Stmt, error) {
	return c.PrepareContext(context.Background(), query)
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
	// Ensures the next transaction will use the session default isolation level if sql.LevelDefault is specified.
	set := "SET @@transaction_isolation = @@SESSION.transaction_isolation"
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
		return nil, fmt.Errorf("Unsupported transaction isolation level (%s)", sql.IsolationLevel(options.Isolation).String())
	}

	// Instead of round trip per sql stmt, use builder to write once.
	b := newBuilderBuffer(c.buf[:0])
	b.WriteExpectOpen(onErrorFail)
	b.WriteStmtExecute(set)
	b.WriteStmtExecute(start)
	b.WriteExpectClose()
	if err := c.sendN(ctx, b.Bytes()); err != nil {
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

func (c *conn) IsValid() bool {
	switch c.status.Get() {
	case statusOK:
		return true
	}
	return false
}

func (c *conn) ResetSession(ctx context.Context) error {
	switch c.status.Get() {
	case statusOK:
		if !c.connector.resetKeepOpen {
			if err := c.send(ctx, msg.SessionReset(c.buf[:0], false)); err != nil {
				return driver.ErrBadConn
			}
			if err := c.authenticate(ctx); err != nil {
				return driver.ErrBadConn
			}
		}
		if err := c.sendN(ctx, c.connector.reset); err != nil {
			return driver.ErrBadConn
		}
		return nil
	default:
		return driver.ErrBadConn
	}
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
			return errors.New("time.Duration overflows mysql TIME (838:59:59)")
		}
		if v < -max {
			return errors.New("time.Duration underflows mysql TIME (-838:59:59)")
		}
	default:
		if _, ok := nv.Value.(msg.AnyAppender); ok {
			return nil
		}
		return fmt.Errorf("unsupported type %T", nv.Value)
	}
	return nil
}

func (c *conn) authenticate(ctx context.Context) error {
	err := c.authenticate2(ctx, c.connector.authentication)
	if err == nil {
		return nil
	}
	if e, ok := errs.IsMySQL(err); !ok || e.Code != errs.ErAccessDeniedError {
		return err
	}
	// Error was Access Denied
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
			return fmt.Errorf("failed reading AuthenticateStart response: %w", err)
		}
		switch t {
		case mysqlx.ServerMessages_NOTICE:
			continue

		case mysqlx.ServerMessages_ERROR:
			return c.handleError(b)

		case mysqlx.ServerMessages_SESS_AUTHENTICATE_OK:
			return nil

		case mysqlx.ServerMessages_SESS_AUTHENTICATE_CONTINUE:
			continuer, ok := starter.(authentication.StartContinuer)
			if !ok {
				return errors.New("unexpected AuthenticateContinue")
			}

			var ac mysqlx_session.AuthenticateContinue
			if err := proto.Unmarshal(b, &ac); err != nil {
				return fmt.Errorf("failed to unmarshal AuthenticateContinue: %w", err)
			}
			if err := c.send(ctx, continuer.Continue(c.buf[:0], c.connector, ac.AuthData)); err != nil {
				return fmt.Errorf("failed sending AuthenticateContinue: %w", err)
			}

			for {
				t, b, err := c.readMessage(ctx)
				if err != nil {
					return fmt.Errorf("failed reading AuthenticateContinue response: %w", err)
				}
				switch t {
				case mysqlx.ServerMessages_NOTICE:
					continue

				case mysqlx.ServerMessages_SESS_AUTHENTICATE_OK:
					return nil

				case mysqlx.ServerMessages_ERROR:
					return c.handleError(b)

				default:
					return fmt.Errorf("unexpected server response to AuthenticateContinue %s", t.String())
				}
			}
		default:
			return fmt.Errorf("unexpected server response to AuthenticateStart %s", t.String())
		}
	}
}

func (c *conn) handleError(b []byte) error {
	err := errs.New(b)
	if e, ok := errs.IsMySQL(err); ok && e.Severity == errs.SeverityFatal {
		c.status.Set(statusBad)
	}
	return err
}
