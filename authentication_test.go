package mysqlx

import (
	"context"
	"database/sql"
	"testing"

	"github.com/renthraysk/mysqlx/authentication/mysql41"
	"github.com/renthraysk/mysqlx/authentication/sha256"
)

/*
	CREATE USER usernative IDENTIFIED WITH 'mysql_native_password' BY 'passwordnative';
	CREATE USER usersha256 IDENTIFIED WITH 'sha256_password' BY 'passwordsha256';
	CREATE USER usersha2 IDENTIFIED WITH 'caching_sha2_password' BY 'passwordsha2';
	FLUSH PRIVILEGES;
*/

func FlushAuthenticationCache(tb testing.TB) {
	tb.Helper()

	db := NewDB(tb)
	defer db.Close()
	_, err := db.ExecContext(context.Background(), "FLUSH PRIVILEGES")
	if err != nil {
		tb.Fatalf("Failed to flush privileges: %s", err)
	}
}

func runAuthenticationTests(t *testing.T) {
	t.Helper()

	tests := map[string]struct {
		network string
		addr    string
		options []Option
	}{
		"tcp-mysql41": {
			"tcp",
			ipAddress,
			[]Option{
				WithAuthentication(mysql41.New()),
				WithUserPassword("usernative", "passwordnative"),
			},
		},

		"tls-mysql41": {
			"tcp",
			ipAddress,
			[]Option{
				WithTLSConfig(TLSInsecureSkipVerify()),
				WithAuthentication(mysql41.New()),
				WithUserPassword("usernative", "passwordnative"),
			},
		},

		"tls-sha2": {
			"tcp",
			ipAddress,
			[]Option{
				WithTLSConfig(TLSInsecureSkipVerify()),
				WithAuthentication(sha256.New()),
				WithUserPassword("usersha2", "passwordsha2"),
			},
		},

		"tls-sha256": {
			"tcp",
			ipAddress,
			[]Option{
				WithTLSConfig(TLSInsecureSkipVerify()),
				WithAuthentication(sha256.New()),
				WithUserPassword("usersha256", "passwordsha256"),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			connector, err := New(test.network, test.addr, test.options...)
			if err != nil {
				t.Fatalf("failed to create connector: %s", err)
			}
			db := sql.OpenDB(connector)
			defer db.Close()
			if err = db.Ping(); err != nil {
				t.Fatalf("failed to ping: %s", err)
			}
		})
	}
}

func TestAuthentication(t *testing.T) {
	FlushAuthenticationCache(t)
	t.Run("empty", runAuthenticationTests)
	t.Run("cached", runAuthenticationTests)
}
