package mysqlx

import (
	"context"
	"crypto/tls"
	"database/sql"
	"testing"
)

const (
	ipAddress   = "127.0.0.1:33060"
	sockAddress = "/var/run/mysqld/mysqlx.sock"
)

var tlsConfig = &tls.Config{InsecureSkipVerify: true}

func NewConnector(tb testing.TB) *Connector {
	tb.Helper()

	connector, err := New("tcp", ipAddress, WithUserPassword("usernative", "passwordnative"), WithDatabase("sakila"))
	if err != nil {
		tb.Fatalf("creating connector failed: %s", err)
	}
	return connector
}

type DB interface {
	ExecContext(ctx context.Context, stmt string, args ...interface{}) (sql.Result, error)

	QueryContext(ctx context.Context, stmt string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, stmt string, args ...interface{}) *sql.Row

	PrepareContext(ctx context.Context, stmt string) (*sql.Stmt, error)

	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

	Close() error
}

func NewDB(tb testing.TB) *sql.DB {
	tb.Helper()
	return sql.OpenDB(NewConnector(tb))
}

func NewDBFatalErrors(tb testing.TB) DB {
	tb.Helper()
	return &DBFatal{NewDB(tb), tb}
}

type DBFatal struct {
	*sql.DB
	tb testing.TB
}

func (d *DBFatal) ExecContext(ctx context.Context, stmt string, args ...interface{}) (sql.Result, error) {
	d.tb.Helper()
	r, err := d.DB.ExecContext(ctx, stmt, args...)
	if err != nil {
		d.tb.Fatalf("Exec failed: %s", err)
	}
	return r, err
}

func (d *DBFatal) QueryContext(ctx context.Context, stmt string, args ...interface{}) (*sql.Rows, error) {
	d.tb.Helper()
	r, err := d.DB.QueryContext(ctx, stmt, args...)
	if err != nil {
		d.tb.Fatalf("Query failed: %s", err)
	}
	return r, err
}

func (d *DBFatal) Prepare(stmt string) (*sql.Stmt, error) {
	d.tb.Helper()
	s, err := d.DB.Prepare(stmt)
	if err != nil {
		d.tb.Fatalf("Prepare failed: %s", err)
	}
	return s, err
}

func (d *DBFatal) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	d.tb.Helper()
	s, err := d.DB.BeginTx(ctx, opts)
	if err != nil {
		d.tb.Fatalf("BeginTx failed: %s", err)
	}
	return s, err
}

func (d *DBFatal) Close() error {
	d.tb.Helper()
	err := d.DB.Close()
	if err != nil {
		d.tb.Fatalf("Close failed: %s", err)
	}
	return err
}
