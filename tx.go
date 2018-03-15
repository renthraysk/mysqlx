package mysqlx

import "database/sql/driver"

type tx struct {
	*conn
}

func (t *tx) Commit() error {
	if t.conn == nil {
		return driver.ErrBadConn
	}
	_, err := t.Exec("COMMIT", nil)
	t.conn = nil
	return err
}

func (t *tx) Rollback() error {
	if t.conn == nil {
		return driver.ErrBadConn
	}
	_, err := t.Exec("ROLLBACK", nil)
	t.conn = nil
	return err
}
