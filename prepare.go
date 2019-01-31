package mysqlx

import (
	"context"
	"database/sql/driver"
	"encoding/binary"
	"hash/fnv"

	"github.com/renthraysk/mysqlx/msg"
)

type preparedStmt struct {
	c    *conn
	stmt string
	id   uint32
}

func (s *preparedStmt) Close() error {
	if s.c == nil {
		return nil
	}
	m := msg.NewDeallocate(s.c.buf[:0], s.id)
	_, err := s.c.execMsg(context.Background(), m)
	s.c = nil
	return err
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
	q, err := msg.NewExecuteNamedArgs(s.c.buf[:0], s.id, args)
	if err != nil {
		return nil, err
	}
	return s.c.queryMsg(ctx, q)
}

func actualStmtPreparer(ctx context.Context, c *conn, stmt string) (driver.Stmt, error) {
	// Attempt to assign the same id for the same stmt
	// @TODO Handle hash collisions.

	h := fnv.New32a()
	h.Write([]byte(stmt))
	id := binary.BigEndian.Uint32(h.Sum(c.buf[:0]))

	if _, err := c.execMsg(ctx, msg.NewPrepare(c.buf[:0], id, stmt)); err != nil {
		return nil, err
	}

	return &preparedStmt{
		c:    c,
		id:   id,
		stmt: stmt,
	}, nil
}
