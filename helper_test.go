package mysqlx

import (
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

func NewDB(tb testing.TB) *sql.DB {
	tb.Helper()

	return sql.OpenDB(NewConnector(tb))
}
