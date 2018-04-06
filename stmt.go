package mysqlx

import (
	"context"
	"database/sql/driver"

	"github.com/renthraysk/mysqlx/msg"
)

type stmtPreparer func(ctx context.Context, c *conn, query string) (driver.Stmt, error)

func noStmtPreparer(ctx context.Context, c *conn, query string) (driver.Stmt, error) {
	return &notPreparedStmt{c, query}, nil
}

type notPreparedStmt struct {
	c     *conn
	query string
}

func (s *notPreparedStmt) Close() error {
	return nil
}

func (s *notPreparedStmt) NumInput() int {
	return -1
}

func (s *notPreparedStmt) Exec(args []driver.Value) (driver.Result, error) {
	m, err := msg.StmtValues(s.c.buf[:0], s.query, args)
	if err != nil {
		return nil, err
	}
	return s.c.execMsg(context.Background(), m)
}

func (s *notPreparedStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	return s.c.ExecContext(ctx, s.query, args)
}

func (s *notPreparedStmt) Query(args []driver.Value) (driver.Rows, error) {
	m, err := msg.StmtValues(s.c.buf[:0], s.query, args)
	if err != nil {
		return nil, err
	}
	return s.c.queryMsg(context.Background(), m)
}

func (s *notPreparedStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	return s.c.QueryContext(ctx, s.query, args)
}
