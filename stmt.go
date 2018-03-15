package mysqlx

import (
	"context"
	"database/sql/driver"
)

type stmtPreparer func(c Conn, query string) (driver.Stmt, error)

func noStmtPreparer(c Conn, query string) (driver.Stmt, error) {
	return &notPreparedStmt{c, query}, nil
}

type notPreparedStmt struct {
	c     Conn
	query string
}

func (s *notPreparedStmt) Close() error {
	return nil
}

func (s *notPreparedStmt) NumInput() int {
	return -1
}

func (s *notPreparedStmt) Exec(args []driver.Value) (driver.Result, error) {
	return s.c.Exec(s.query, args)
}

func (s *notPreparedStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	return s.c.ExecContext(ctx, s.query, args)
}

func (s *notPreparedStmt) Query(args []driver.Value) (driver.Rows, error) {
	return s.c.Query(s.query, args)
}

func (s *notPreparedStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	return s.c.QueryContext(ctx, s.query, args)
}
