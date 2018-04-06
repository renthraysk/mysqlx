package mysqlx

import (
	"context"
	"database/sql"
	"testing"
)

func TestBeginTx(t *testing.T) {
	t.Skipf("Can not determine current transaction's isolation level: https://bugs.mysql.com/bug.php?id=53341")

	isos := map[sql.IsolationLevel]string{
		sql.LevelReadUncommitted: "READ-UNCOMMITTED",
		sql.LevelReadCommitted:   "READ-COMMITTED",
		sql.LevelRepeatableRead:  "REPEATABLE READ",
		sql.LevelSerializable:    "SERIALIZABLE",
	}

	db := NewDBFatalErrors(t)
	defer db.Close()

	for level, name := range isos {
		tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: level})
		if err != nil {
			t.Fatalf("BeginTx failed: %s for %s", err, name)
		}
		_, err = tx.QueryContext(context.Background(), "SELECT @@session.transaction_isolation")
		if err != nil {
			t.Fatalf("SELECT @@transaction_isolation failed: %s", err)
		}
		tx.Rollback()
	}
}
