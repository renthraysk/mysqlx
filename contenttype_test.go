package mysqlx

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)

	var b []byte
	var retrieved Thing

	r, err := db.ExecContext(context.Background(), "INSERT INTO json(json) VALUES(?)", JSON(j))
	assert.NoError(t, err)
	n, err := r.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), n)
	id, err := r.LastInsertId()
	assert.NoError(t, err)
	assert.NoError(t, db.QueryRowContext(context.Background(), "SELECT json FROM json WHERE id = ?", id).Scan(&b))
	assert.NoError(t, json.Unmarshal(b, &retrieved))
	assert.Equal(t, expected, &retrieved)

}

func TestSendJSONNull(t *testing.T) {
	var b []byte

	db := NewDB(t)
	defer db.Close()

	r, err := db.ExecContext(context.Background(), "INSERT INTO json(json) VALUES(?)", JSON(nil))
	assert.NoError(t, err)
	n, err := r.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), n)
	id, err := r.LastInsertId()
	assert.NoError(t, err)
	assert.NoError(t, db.QueryRowContext(context.Background(), "SELECT json FROM json WHERE id = ?", id).Scan(&b))
	assert.Nil(t, b)
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
	var retrieved Thing

	x, err := xml.Marshal(expected)
	assert.NoError(t, err)
	r, err := db.ExecContext(context.Background(), "INSERT INTO xml(xml) VALUES(?)", XML(x))
	assert.NoError(t, err)
	n, err := r.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), n)
	id, err := r.LastInsertId()
	assert.NoError(t, err)
	assert.NoError(t, db.QueryRowContext(context.Background(), "SELECT xml FROM xml WHERE id = ?", id).Scan(&b))
	assert.NoError(t, xml.Unmarshal(b, &retrieved))
	assert.Equal(t, expected, &retrieved)
}
