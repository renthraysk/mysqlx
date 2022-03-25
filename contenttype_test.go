package mysqlx

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"testing"
)

func TestSendJSON(t *testing.T) {

	db := NewDB(t)
	defer db.Close()

	type Thing struct {
		Name   string `json:"name,omitempty"`
		Number int    `json:"number,omitempty"`
	}

	expected := &Thing{
		Name:   "Alice",
		Number: 42,
	}

	j, err := json.Marshal(expected)
	assertNoError(t, err)

	var b []byte
	var actual Thing

	r, err := db.ExecContext(context.Background(), "INSERT INTO json(json) VALUES(?)", JSON(j))
	assertNoError(t, err)
	n, err := r.RowsAffected()
	assertNoError(t, err)
	assertComparableEqual(t, int64(1), n)
	id, err := r.LastInsertId()
	assertNoError(t, err)
	assertNoError(t, db.QueryRowContext(context.Background(), "SELECT json FROM json WHERE id = ?", id).Scan(&b))
	assertNoError(t, json.Unmarshal(b, &actual))

	assertComparableEqual(t, expected.Name, actual.Name)
	assertComparableEqual(t, expected.Number, actual.Number)
}

func TestSendJSONNull(t *testing.T) {
	var b []byte

	db := NewDB(t)
	defer db.Close()

	r, err := db.ExecContext(context.Background(), "INSERT INTO json(json) VALUES(?)", JSON(nil))
	assertNoError(t, err)
	n, err := r.RowsAffected()
	assertNoError(t, err)
	assertComparableEqual(t, int64(1), n)
	id, err := r.LastInsertId()
	assertNoError(t, err)
	assertNoError(t, db.QueryRowContext(context.Background(), "SELECT json FROM json WHERE id = ?", id).Scan(&b))

	if !isNil(b) {
		t.Fatalf("expected nil, got %T(%v)", b, b)
	}
}

func TestSendXML(t *testing.T) {

	db := NewDB(t)
	defer db.Close()

	type Thing struct {
		XMLName xml.Name
		Name    string `xml:"name"`
		Number  int    `xml:"number"`
	}

	expected := &Thing{
		XMLName: xml.Name{Local: "thing"},
		Name:    "Alice",
		Number:  42,
	}
	var b []byte
	var actual Thing

	x, err := xml.Marshal(expected)
	assertNoError(t, err)
	r, err := db.ExecContext(context.Background(), "INSERT INTO xml(xml) VALUES(?)", XML(x))
	assertNoError(t, err)
	n, err := r.RowsAffected()
	assertNoError(t, err)
	assertComparableEqual(t, int64(1), n)
	id, err := r.LastInsertId()
	assertNoError(t, err)
	assertNoError(t, db.QueryRowContext(context.Background(), "SELECT xml FROM xml WHERE id = ?", id).Scan(&b))
	assertNoError(t, xml.Unmarshal(b, &actual))

	assertComparableEqual(t, expected.Name, actual.Name)
	assertComparableEqual(t, expected.Number, actual.Number)
}
