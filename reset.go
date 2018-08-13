package mysqlx

import (
	"context"
	"database/sql/driver"

	"github.com/renthraysk/mysqlx/msg"
)

// SessionResetter function to call to reset the connection for reuse.
type SessionResetter func(ctx context.Context, c *conn) error

// NoSessionResetter a no op session resetter, historically equivalent behaviour
func NoSessionResetter(ctx context.Context, c *conn) error {
	return nil
}

// HardSessionResetter a full connection reset. Transactions closed, prepared statements deleted, temporary tables dropped and session variables reset to global defaults
func HardSessionResetter(ctx context.Context, c *conn) error {
	if err := c.send(ctx, msg.SessionReset(c.buf[:0])); err != nil {
		return driver.ErrBadConn
	}
	c.openTxCount = 0
	if err := c.authenticate(ctx); err != nil {
		return driver.ErrBadConn
	}
	return nil
}
