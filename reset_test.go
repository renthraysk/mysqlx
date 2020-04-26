package mysqlx

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/renthraysk/mysqlx/authentication/mysql41"
	"github.com/renthraysk/mysqlx/errs"
)

func DB(t *testing.T, opt Option, q string) (string, error) {
	connector, err := New("tcp", "127.0.0.1:33060",
		WithAuthentication(mysql41.New()),
		WithUserPassword("usernative", "passwordnative"),
		opt,
	)
	if err != nil {
		return "", err
	}
	db := sql.OpenDB(connector)
	defer db.Close()

	if err := db.Ping(); err != nil {
		return "", err
	}

	var v string

	err = db.QueryRow(q).Scan(&v)
	return v, err
}

func DBSQLMode(t *testing.T, mode string) (string, error) {
	return DB(t, WithSQLMode(mode), "SELECT @@sql_mode")
}

func TestInitSQLMode(t *testing.T) {
	// Test setting @@sql_mode twice with two difference values, incase the mysql server default is one or the other.

	{
		const expected = "REAL_AS_FLOAT,PIPES_AS_CONCAT,ANSI_QUOTES,IGNORE_SPACE,ONLY_FULL_GROUP_BY,ANSI"

		got, err := DBSQLMode(t, "ANSI")
		if err != nil {
			t.Fatalf("failed to set @@sql_mode: %s", err)
		}
		if got != expected {
			t.Fatalf("failed to set @@sql_mode to ANSI, expected %s got %s", expected, got)
		}
	}
	{
		const expected = "STRICT_TRANS_TABLES,STRICT_ALL_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,TRADITIONAL,NO_ENGINE_SUBSTITUTION"

		got, err := DBSQLMode(t, "TRADITIONAL")
		if err != nil {
			t.Fatalf("failed to set @@sql_mode: %s", err)
		}
		if got != expected {
			t.Fatalf("failed to set @@sql_mode to ANSI, expected %s got %s", expected, got)
		}
	}
}

func TestInitSQLModeError(t *testing.T) {
	_, err := DBSQLMode(t, "NONSENSE")
	if err == nil {
		t.Fatalf("failed to trigger error setting @@sql_mode to an invalid value")
	}

	var ers errs.Errors

	if errors.As(err, &ers) {
		t.Fatalf("expected mysqlx.Errors")
	}
	// Error happens on the third statement... first is an expect open, 2nd is session-reset(keepOpen=true) or just ping (pre mysql 8.0.16)
	if e, ok := ers[2].(*errs.Error); !ok ||
		e.Code != errs.ErWrongValueForVar ||
		e.Msg != "Variable 'sql_mode' can't be set to the value of 'NONSENSE'" {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestInitSessionVars(t *testing.T) {
	{
		got, err := DB(t, WithSessionVars(SessionVars{
			"group_concat_max_len": 1234,
		}), "SELECT @@group_concat_max_len")

		if err != nil {
			t.Fatalf("failed to set group_concat_max_len: %s", err)
		}
		if got != "1234" {
			t.Fatalf("failed to set group_concat_max_len, expected %s, got %s", "1234", got)
		}
	}
	{
		got, err := DB(t, WithSessionVars(SessionVars{
			"group_concat_max_len": 12345,
		}), "SELECT @@group_concat_max_len")

		if err != nil {
			t.Fatalf("failed to set group_concat_max_len: %s", err)
		}
		if got != "12345" {
			t.Fatalf("failed to set group_concat_max_len, expected %s, got %s", "12345", got)
		}
	}
}

func TestDefaultTxIso(t *testing.T) {
	{
		got, err := DB(t, WithDefaultTxIsolation(sql.LevelReadUncommitted), "SELECT @@SESSION.transaction_isolation")
		if err != nil {
			t.Fatalf("failed to set session isolation level: %s", err)
		}
		if got != "READ-UNCOMMITTED" {
			t.Fatalf("failed to set session isolation level, expected %s, got %s", "READ-UNCOMMITTED", got)
		}
	}
	{
		got, err := DB(t, WithDefaultTxIsolation(sql.LevelReadCommitted), "SELECT @@SESSION.transaction_isolation")
		if err != nil {
			t.Fatalf("failed to set session isolation level: %s", err)
		}
		if got != "READ-COMMITTED" {
			t.Fatalf("failed to set session isolation level, expected %s, got %s", "READ-COMMITTED", got)
		}
	}
}
