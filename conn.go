package mysqlx

import (
	"context"
	"crypto/tls"
	"database/sql"
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/renthraysk/mysqlx/authentication"
	"github.com/renthraysk/mysqlx/authentication/plain"
	"github.com/renthraysk/mysqlx/msg"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_session"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type execMsgExecuter interface {
	execMsg(ctx context.Context, msg msg.Msg) (driver.Result, error)
}

type queryMsgExecuter interface {
	queryMsg(ctx context.Context, msg msg.Msg) (driver.Rows, error)
}

// Conn the interface of our driver.Conn plus extra methods
type Conn interface {
	driver.Conn

	driver.Execer
	driver.Queryer

	driver.QueryerContext
	driver.ExecerContext
	driver.ConnBeginTx
	driver.Pinger

	driver.NamedValueChecker
	driver.SessionResetter

	execMsgExecuter
	queryMsgExecuter
}

type conn struct {
	netConn   net.Conn
	connector *Connector

	buf    []byte
	offset uint
	length uint
}

func (c *conn) read(ctx context.Context, n uint) ([]byte, error) {
	if c.length < n {
		if b := c.buf[c.offset+c.length:]; uint(len(b)) < n {
			if uint(len(c.buf)) < n {
				b := make([]byte, n)
				copy(b, c.buf[c.offset:c.offset+c.length])
				c.buf = b
			} else {
				copy(c.buf, c.buf[c.offset:c.offset+c.length])
			}
			c.offset = 0
		}

		deadline, _ := ctx.Deadline()
		if err := c.netConn.SetReadDeadline(deadline); err != nil {
			return nil, err
		}
		for c.length < n {
			nn, err := c.netConn.Read(c.buf[c.offset+c.length:])
			c.length += uint(nn)
			if err != nil {
				return nil, err
			}
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
			}
		}
	}

	c.offset += n
	c.length -= n

	return c.buf[c.offset-n : c.offset : c.offset], nil
}

func (c *conn) readMessage(ctx context.Context) (mysqlx.ServerMessages_Type, []byte, error) {
	b, err := c.read(ctx, 5)
	if err != nil {
		return 0, nil, err
	}
	t := mysqlx.ServerMessages_Type(b[4])
	n := uint(binary.LittleEndian.Uint32(b))
	if n <= 1 {
		return t, nil, nil
	}
	b, err = c.read(ctx, n-1)
	return t, b, err
}

func (c *conn) send(ctx context.Context, msg msg.Msg) error {
	c.offset = 0
	c.length = 0
	deadline, _ := ctx.Deadline()
	if err := c.netConn.SetWriteDeadline(deadline); err != nil {
		return errors.Wrap(err, "unable to set deadline")
	}
	_, err := msg.WriteTo(c.netConn)
	return err
}

func (c *conn) execMsg(ctx context.Context, m msg.Msg) (driver.Result, error) {
	err := c.send(ctx, m)
	if err != nil {
		return nil, err
	}
readExecResponse:
	t, b, err := c.readMessage(ctx)
	if err != nil {
		return nil, err
	}
	switch t {
	case mysqlx.ServerMessages_OK:
		return nil, nil

	case mysqlx.ServerMessages_ERROR:
		return nil, newError(b)

	case mysqlx.ServerMessages_SQL_STMT_EXECUTE_OK:
		break

	case mysqlx.ServerMessages_NOTICE:
		goto readExecResponse
	default:
		goto readExecResponse
	}
	return nil, nil
}

func (c *conn) Exec(stmt string, args []driver.Value) (driver.Result, error) {
	s, err := msg.StmtValues(c.buf[:0], stmt, args)
	if err != nil {
		return nil, err
	}
	return c.execMsg(context.Background(), s)
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

readColumnMetaData:
	r.last.t, r.last.b, r.last.err = c.readMessage(ctx)
	if r.last.err != nil {
		return nil, errors.Wrap(r.last.err, "failed to read query response")
	}
	switch r.last.t {

	case mysqlx.ServerMessages_ERROR:
		r.last.err = newError(r.last.b)
		return nil, r.last.err

	case mysqlx.ServerMessages_SQL_STMT_EXECUTE_OK:
		return r, nil

	case mysqlx.ServerMessages_RESULTSET_FETCH_DONE, mysqlx.ServerMessages_RESULTSET_FETCH_DONE_MORE_RESULTSETS:
		return r, nil

	case mysqlx.ServerMessages_RESULTSET_COLUMN_META_DATA:

		var cmd *columnMetaData

		n := len(r.columns)
		if n < len(r.buf) {
			cmd = &r.buf[n]
		} else {
			cmd = new(columnMetaData)
		}
		if err := cmd.unmarshal(r.last.b); err != nil {
			return nil, err
		}
		r.columns = append(r.columns, cmd)

		goto readColumnMetaData
	}

	return r, nil
}

func (c *conn) Query(stmt string, args []driver.Value) (driver.Rows, error) {
	s, err := msg.StmtValues(c.buf[:0], stmt, args)
	if err != nil {
		return nil, err
	}
	return c.queryMsg(context.Background(), s)
}

func (c *conn) QueryContext(ctx context.Context, stmt string, args []driver.NamedValue) (driver.Rows, error) {
	s, err := msg.StmtNamedValues(c.buf[:0], stmt, args)
	if err != nil {
		return nil, err
	}
	return c.queryMsg(ctx, s)
}

func (c *conn) Prepare(query string) (driver.Stmt, error) {
	return c.connector.stmtPreparer(c, query)
}

func (c *conn) Close() error {
	if c.netConn == nil {
		return nil
	}
	ctx := context.Background()

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
		return nil, fmt.Errorf("Unsupported transaction isolation level (%d)", options.Isolation)
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

func (c *conn) Begin() (driver.Tx, error) {
	if _, err := c.Exec("START TRANSACTION", nil); err != nil {
		return nil, err
	}
	return &tx{c}, nil
}

func (c *conn) Ping(ctx context.Context) error {
	ping := msg.NewStmtExecute(c.buf[:0], "ping")
	ping.SetNamespace("mysqlx")
	_, err := c.execMsg(ctx, ping)
	return err
}

func (c *conn) ResetSession(ctx context.Context) error {
	return c.connector.sessionResetter(ctx, c)
}

func (c *conn) CheckNamedValue(nv *driver.NamedValue) error {
	if nv.Value == nil {
		return nil
	}
	switch nv.Value.(type) {
	case uint64, int64, string, []byte, float32, float64, bool, time.Time:
		// Protocol supported types.
	case uint8, uint16, uint32, uint, int8, int16, int32, int:
		// Supported via conversion to a type in above case.
	default:
		return fmt.Errorf("Unsupported type %T", nv.Value)
	}
	return nil
}

func (c *conn) authenticate(ctx context.Context) error {
	err := c.authenticate2(ctx, c.connector.authentication)
	if err == nil {
		return nil
	}
	if e, ok := err.(*Error); !ok || e.Code != 1045 { // Invalid user or password (code 1045)
		return err
	}
	switch c.netConn.(type) {
	case *tls.Conn, *net.UnixConn:
		// Connected securely, so can attempt to authenticate with PLAIN
		if err2 := c.authenticate2(ctx, plain.New()); err2 == nil {
			return nil
		}
	default:
		// @TODO Need decide what to do here..
		// https://dev.mysql.com/doc/refman/8.0/en/x-plugin-sha2-cache-plugin.html
		// Current feeling is to not allow authentication with sha2 over non secure connections.
	}
	return err
}

func (c *conn) authenticate2(ctx context.Context, starter authentication.Starter) error {
	if err := c.send(ctx, starter.Start(c.buf[:0], c.connector)); err != nil {
		return err
	}

readAuthenticateStartResponse:
	t, b, err := c.readMessage(ctx)
	if err != nil {
		return errors.Wrap(err, "failed reading AuthenticateStart response")
	}
	switch t {

	case mysqlx.ServerMessages_ERROR:
		return newError(b)

	case mysqlx.ServerMessages_SESS_AUTHENTICATE_OK:
		return nil

	case mysqlx.ServerMessages_NOTICE:
		goto readAuthenticateStartResponse

	case mysqlx.ServerMessages_SESS_AUTHENTICATE_CONTINUE:
		continuer, ok := starter.(authentication.StartContinuer)
		if !ok {
			return errors.New("unexpected AuthenticateContinue")
		}

		var ac mysqlx_session.AuthenticateContinue
		if err := proto.Unmarshal(b, &ac); err != nil {
			return errors.Wrap(err, "failed to unmarhsal AuthenticateContinue")
		}
		if err := c.send(ctx, continuer.Continue(c.buf[:0], c.connector, ac.GetAuthData())); err != nil {
			return errors.Wrap(err, "failed sending AuthenticateContinue")
		}
	readAuthenticateContinueResponse:
		t, b, err := c.readMessage(ctx)
		if err != nil {
			return errors.Wrap(err, "failed reading AuthenticateContinue response")
		}
		switch t {
		case mysqlx.ServerMessages_ERROR:
			return newError(b)

		case mysqlx.ServerMessages_NOTICE:
			goto readAuthenticateContinueResponse

		case mysqlx.ServerMessages_SESS_AUTHENTICATE_OK:
			return nil

		default:
			return fmt.Errorf("unexpected server response to AuthenticateContinue %s(%d)", t.String(), t)
		}
	default:
	}
	return fmt.Errorf("unexpected server response to AuthenticateStart %s(%d)", t.String(), t)
}
