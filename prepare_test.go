package mysqlx

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepare(t *testing.T) {
	const (
		SQLTEXT = "SELECT ?"
	)
	ctx := context.Background()

	connector := NewConnector(t)

	db := sql.OpenDB(connector)
	defer db.Close()

	// Grab single connection
	cnn, err := db.Conn(ctx)
	assert.NoError(t, err)

	stmt, err := cnn.PrepareContext(ctx, SQLTEXT)
	assert.NoError(t, err)

	// Make sure our SQL didn't get mangled...
	rows, err := cnn.QueryContext(ctx, "SELECT 1 from performance_schema.prepared_statements_instances WHERE SQL_TEXT = ?", SQLTEXT)
	assert.NoError(t, err)
	assert.True(t, rows.Next())
	assert.False(t, rows.Next())
	assert.NoError(t, rows.Close())

	values := []interface{}{
		nil,
		0,
		1,
		1.5,
		42,
		"abc",
		[]byte{'x', 'y', 'z'},
	}

	for _, v := range values {
		t.Run(fmt.Sprintf("test.%v", v), func(t *testing.T) {

			var r interface{}

			rows, err := stmt.Query(v)
			assert.NoError(t, err)
			assert.True(t, rows.Next())
			assert.NoError(t, rows.Scan(&r))
			assert.EqualValues(t, v, r)
			assert.NoError(t, rows.Close())
		})
	}

	assert.NoError(t, stmt.Close())
	assert.NoError(t, cnn.Close())
	assert.NoError(t, db.Close())
}