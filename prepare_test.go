package mysqlx

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepare(t *testing.T) {
	connector := NewConnector(t)

	db := sql.OpenDB(connector)
	defer db.Close()
	stmt, err := db.Prepare("SELECT ?")
	if err != nil {
		t.Fatal(err)
	}
	defer stmt.Close()

	values := []interface{}{
		false,
		true,
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
}
