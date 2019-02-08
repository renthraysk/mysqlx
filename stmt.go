package mysqlx

import (
	"context"
	"database/sql/driver"

	"github.com/renthraysk/mysqlx/msg"
)

type stmt struct {
	c  *conn
	id uint32
}

func (s *stmt) Close() error {
	if s.c == nil {
		return nil
	}
	m := msg.NewDeallocate(s.c.buf[:0], s.id)
	_, err := s.c.execMsg(context.Background(), m)
	s.c = nil
	return err
}

func (s *stmt) NumInput() int {
	return -1
}

// Exec forced deprecated implementation by database/sql Stmt interface
func (s *stmt) Exec(args []driver.Value) (driver.Result, error) {
	e, err := msg.NewExecuteArgs(s.c.buf[:0], s.id, args)
	if err != nil {
		return nil, err
	}
	return s.c.execMsg(context.Background(), e)
}

func (s *stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	e, err := msg.NewExecuteNamedArgs(s.c.buf[:0], s.id, args)
	if err != nil {
		return nil, err
	}
	return s.c.execMsg(ctx, e)
}

// Query forced deprecated implementation by database/sql Stmt interface
func (s *stmt) Query(args []driver.Value) (driver.Rows, error) {
	e, err := msg.NewExecuteArgs(s.c.buf[:0], s.id, args)
	if err != nil {
		return nil, err
	}
	return s.c.queryMsg(context.Background(), e)
}

func (s *stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	q, err := msg.NewExecuteNamedArgs(s.c.buf[:0], s.id, args)
	if err != nil {
		return nil, err
	}
	return s.c.queryMsg(ctx, q)
}
