package mysqlx

import (
	"context"
	"database/sql/driver"

	"github.com/renthraysk/mysqlx/msg"
)

type sessionResetter func(ctx context.Context, c *conn) error

func noSessionResetter(ctx context.Context, c *conn) error {
	return nil
}

func hardSessionResetter(ctx context.Context, c *conn) error {
	if err := c.send(ctx, msg.SessionReset(c.buf[:0])); err != nil {
		return driver.ErrBadConn
	}
	if err := c.authenticate(ctx); err != nil {
		return driver.ErrBadConn
	}
	return nil
}
