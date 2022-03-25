package mysqlx

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
)

const (
	ipAddress   = "127.0.0.1:33060"
	sockAddress = "/var/run/mysqld/mysqlx.sock"
)

func NewConnector(tb testing.TB) *Connector {
	tb.Helper()

	connector, err := New("tcp", ipAddress,
		WithUserPassword("usernative", "passwordnative"),
		WithDatabase("gotest"),
		WithDefaultConnectAttrs(),
	)
	if err != nil {
		tb.Fatalf("failed creating connector: %s", err)
	}
	return connector
}

func NewDB(tb testing.TB) *sql.DB {
	tb.Helper()
	return sql.OpenDB(NewConnector(tb))
}

func query(tb testing.TB, sql string, args []any, scan func(rows *sql.Rows) error) {
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

func isNil(value any) bool {
	const Nilable = 1<<reflect.Chan | 1<<reflect.Func | 1<<reflect.Interface | 1<<reflect.Map | 1<<reflect.Ptr | 1<<reflect.Slice

	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	return (1<<v.Kind())&Nilable != 0 && v.IsNil()
}

func assertComparableEqual[T comparable](t *testing.T, expected T, actual any) {
	t.Helper()
	if v, ok := actual.(T); !ok || v != expected {
		t.Fatalf("expected %T(%v), got %T(%v)", expected, expected, actual, actual)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
}
