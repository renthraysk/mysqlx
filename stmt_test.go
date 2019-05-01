package mysqlx

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	require.NoError(t, err)

	stmt, err := cnn.PrepareContext(ctx, SQLTEXT)
	require.NoError(t, err)

	// Make sure our SQL didn't get mangled...
	rows, err := cnn.QueryContext(ctx, "SELECT 1 from performance_schema.prepared_statements_instances WHERE SQL_TEXT = ?", SQLTEXT)
	require.NoError(t, err)

	require.True(t, rows.Next())
	require.False(t, rows.Next())
	require.NoError(t, rows.Close())

	values := []interface{}{
		nil,
		0,
		1,
		float64(1.5),
		float32(1.34),
		42,
		"abc",
		[]byte{'x', 'y', 'z'},
		JSON([]byte("{}")),
		XML([]byte("<foo />")),
		true,
		false,
	}

	for _, v := range values {
		t.Run(fmt.Sprintf("test.%T(%v)", v, v), func(t *testing.T) {

			var r interface{}

			rows, err := stmt.Query(v)
			require.NoError(t, err)
			require.True(t, rows.Next())
			require.NoError(t, rows.Scan(&r))
			switch vv := v.(type) {
			case bool:
				if vv {
					// nonzero is true
					require.False(t, assert.ObjectsAreEqual(0, r))
				} else {
					require.EqualValues(t, 0, r)
				}
			default:
				require.EqualValues(t, v, r)
			}
			require.NoError(t, rows.Close())
		})
	}

	require.NoError(t, stmt.Close())
	require.NoError(t, cnn.Close())
	require.NoError(t, db.Close())
}
