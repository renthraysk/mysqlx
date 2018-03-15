package mysqlx

import (
	"context"

	"github.com/renthraysk/mysqlx/msg"

	"github.com/pkg/errors"
)

type sessionResetter func(ctx context.Context, c *conn) error

func noSessionResetter(ctx context.Context, c *conn) error {
	return nil
}

func hardSessionResetter(ctx context.Context, c *conn) error {
	if err := c.send(ctx, msg.SessionReset(c.buf[:0])); err != nil {
		return errors.Wrap(err, "failed to reset")
	}
	return c.authenticate(ctx, c.connector.authentication)
}
