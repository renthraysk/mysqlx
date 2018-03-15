package mysqlx

import (
	"crypto/tls"
	"database/sql"
	"testing"

	"github.com/renthraysk/mysqlx/authentication/native"
	"github.com/renthraysk/mysqlx/authentication/sha256"
)

func TestAuthentication(t *testing.T) {

	/*
		CREATE USER usernative IDENTIFIED WITH 'mysql_native_password' BY 'passwordnative';
		CREATE USER usersha256 IDENTIFIED WITH 'sha256_password' BY 'passwordsha256';
		CREATE USER usersha2 IDENTIFIED WITH 'caching_sha2_password' BY 'passwordsha2';
		FLUSH PRIVILEGES;
	*/

	tests := map[string]struct {
		network string
		addr    string
		options []Option
	}{
		"tcp-native": {
			"tcp",
			ipAddress,
			[]Option{WithAuthentication(native.New()), WithUserPassword("usernative", "passwordnative")},
		},

		"tls-native": {
			"tcp",
			ipAddress,
			[]Option{WithTLSConfig(tlsConfig), WithAuthentication(native.New()), WithUserPassword("usernative", "passwordnative")},
		},

		"tls-sha2": {
			"tcp",
			ipAddress,
			[]Option{WithTLSConfig(tlsConfig), WithAuthentication(sha256.New()), WithUserPassword("usersha2", "passwordsha2")},
		},

		"tls-sha256": {
			"tcp",
			ipAddress,
			[]Option{WithTLSConfig(tlsConfig), WithAuthentication(sha256.New()), WithUserPassword("usersha256", "passwordsha256")},
		},
	}

	//FlushAuthenticationCache(t)

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
