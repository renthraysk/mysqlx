package mysqlx

import (
	"context"
	"database/sql"
	"testing"
	"time"
)

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
	specialFeatures    Set
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

func TestSakilaFilmQuery(t *testing.T) {
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
