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

	connector, err := New("tcp", ipAddress,
		WithUserPassword("usernative", "passwordnative"))
	if err != nil {
		tb.Fatalf("creating connector failed: %s", err)
	}
	return connector
}

func NewDB(tb testing.TB) *sql.DB {
	tb.Helper()
	return sql.OpenDB(NewConnector(tb))
}

func query(tb testing.TB, sql string, args []interface{}, scan func(rows *sql.Rows) error) {
	tb.Helper()

	db := NewDB(tb)
	defer db.Close()
	rows, err := db.QueryContext(context.Background(), sql, args...)
	if err != nil {
		tb.Fatalf("QueryContext failed: %s", err)
	}
	defer rows.Close()
	if !rows.Next() {
		tb.Fatalf("no row returned: %s", rows.Err())
	}
	if err := scan(rows); err != nil {
		tb.Fatalf("scan failed: %s", err)
	}
	if rows.Next() {
		tb.Fatal("more than one row returned")
	}
	if err := rows.Err(); err != nil {
		tb.Fatalf("rows.Err: %s", err)
	}
	if err := rows.Close(); err != nil {
		tb.Fatalf("Close return error: %s", err)
	}
}
