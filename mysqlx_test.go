package mysqlx

import (
	"context"
	"database/sql"
	"testing"
	"time"
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

const SelectAll = `SELECT 
	film_id,
    title,
    description,
    release_year,
    language_id,
    original_language_id,
    rental_duration,
    rental_rate,
    length,
    replacement_cost,
    rating,
    special_features,
    last_update
FROM film`

type film struct {
	filmID             uint16
	title              string
	description        sql.NullString
	releaseYear        sql.NullInt64
	languageID         uint8
	originalLanguageID sql.NullInt64
	rentalDuration     uint8
	rentalRate         string
	length             sql.NullInt64
	replacementCost    string
	rating             sql.NullString
	specialFeatures    sql.NullString
	lastUpdate         time.Time
}

func (f *film) Scan(rows *sql.Rows) error {
	return rows.Scan(
		&f.filmID,
		&f.title,
		&f.description,
		&f.releaseYear,
		&f.languageID,
		&f.originalLanguageID,
		&f.rentalDuration,
		&f.rentalRate,
		&f.length,
		&f.replacementCost,
		&f.rating,
		&f.specialFeatures,
		&f.lastUpdate,
	)
}

func TestQuery(t *testing.T) {
	var f film

	db := NewDBFatalErrors(t)
	defer db.Close()

	rows, _ := db.QueryContext(context.Background(), SelectAll)
	for rows.Next() {
		if err := f.Scan(rows); err != nil {
			t.Fatalf("scan failed: %s", err)
		}
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("rows error: %s", err)
	}

	rows.Close()
}
