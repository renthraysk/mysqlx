package mysqlx

import (
	"context"
	"database/sql/driver"

	"github.com/pkg/errors"

	"github.com/renthraysk/mysqlx/msg"
)

type preparedStmt struct {
	c  *conn
	id uint32
}

func (s *preparedStmt) Close() error {
	return nil
}

func (s *preparedStmt) NumInput() int {
	return -1
}

// Exec forced deprecated implementation by database/sql Stmt interface
func (s *preparedStmt) Exec(args []driver.Value) (driver.Result, error) {
	e, err := msg.NewExecuteArgs(s.c.buf[:0], s.id, args)
	if err != nil {
		return nil, err
	}
	return s.c.execMsg(context.Background(), e)
}

func (s *preparedStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	e, err := msg.NewExecuteNamedArgs(s.c.buf[:0], s.id, args)
	if err != nil {
		return nil, err
	}
	return s.c.execMsg(ctx, e)
}

// Query forced deprecated implementation by database/sql Stmt interface
func (s *preparedStmt) Query(args []driver.Value) (driver.Rows, error) {
	e, err := msg.NewExecuteArgs(s.c.buf[:0], s.id, args)
	if err != nil {
		return nil, err
	}
	return s.c.queryMsg(context.Background(), e)
}

func (s *preparedStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	e, err := msg.NewExecuteNamedArgs(s.c.buf[:0], s.id, args)
	if err != nil {
		return nil, err
	}
	return s.c.queryMsg(ctx, e)
}

func actualStmtPreparer(ctx context.Context, c *conn, query string) (driver.Stmt, error) {
	var id uint32

	_, err := c.execMsg(ctx, msg.NewPrepare(c.buf[:0], id, query))
	if err != nil {
		return nil, errors.Wrap(err, "prepare failed")
	}

	return &preparedStmt{c, id}, nil
}
