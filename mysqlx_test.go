package mysqlx

import (
	"context"
	"database/sql"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNull(t *testing.T) {
	var null interface{}

	null = 42
	query(t, "SELECT ?", []interface{}{nil}, func(rows *sql.Rows) error { return rows.Scan(&null) })
	assert.Nil(t, null)
}

func TestBool(t *testing.T) {
	expected := []bool{false, true}
	in := make([]interface{}, len(expected))
	for i := 0; i < len(expected); i++ {
		in[i] = interface{}(expected[i])
	}
	out := make([]bool, len(expected))

	query(t, "SELECT ?, ?", in, func(rows *sql.Rows) error { return rows.Scan(&out[0], &out[1]) })
	assert.Equal(t, expected, out)
}

func TestUint(t *testing.T) {
	expected := []uint64{0, math.MaxUint64}
	in := make([]interface{}, len(expected))
	for i := 0; i < len(expected); i++ {
		in[i] = interface{}(expected[i])
	}
	out := make([]uint64, len(expected))

	query(t, "SELECT ?, ?", in, func(rows *sql.Rows) error { return rows.Scan(&out[0], &out[1]) })
	assert.Equal(t, expected, out)
}

func TestInt(t *testing.T) {
	expected := []int64{math.MinInt64, 0, math.MaxInt64}
	in := make([]interface{}, len(expected))
	for i := 0; i < len(expected); i++ {
		in[i] = interface{}(expected[i])
	}
	out := make([]int64, len(expected))

	query(t, "SELECT ?, ?, ?", in, func(rows *sql.Rows) error { return rows.Scan(&out[0], &out[1], &out[2]) })
	assert.Equal(t, expected, out)
}

func TestFloat32(t *testing.T) {
	// @TODO math.MaxFloat32 appears to get truncated on a roundtrip
	expected := []float32{0, math.SmallestNonzeroFloat32, math.MaxFloat32 - 3.5e+32}
	in := make([]interface{}, len(expected))
	for i := 0; i < len(expected); i++ {
		in[i] = interface{}(expected[i])
	}
	out := make([]float32, len(expected))

	query(t, "SELECT ?, ?, ?", in, func(rows *sql.Rows) error { return rows.Scan(&out[0], &out[1], &out[2]) })
	assert.Equal(t, expected, out)
}

func TestFloat64(t *testing.T) {
	// @TODO math.FloaMaxFloat64 appears to get truncated on a roundtrip
	expected := []float64{0, math.SmallestNonzeroFloat64, math.MaxFloat64 - 3.1348623157e+302}
	in := make([]interface{}, len(expected))
	for i := 0; i < len(expected); i++ {
		in[i] = interface{}(expected[i])
	}
	out := make([]float64, len(expected))

	query(t, "SELECT ?, ?, ?", in, func(rows *sql.Rows) error { return rows.Scan(&out[0], &out[1], &out[2]) })
	assert.Equal(t, expected, out)
}

func TestString(t *testing.T) {
	expected := []string{"", "abc", "abcdef"}
	in := make([]interface{}, len(expected))
	for i := 0; i < len(expected); i++ {
		in[i] = interface{}(expected[i])
	}
	out := make([]string, len(expected))

	query(t, "SELECT ?, ?, ?", in, func(rows *sql.Rows) error { return rows.Scan(&out[0], &out[1], &out[2]) })
	assert.Equal(t, expected, out)
}

func TestBytes(t *testing.T) {
	expected := [][]byte{[]byte{}, []byte{0x00}, []byte("abcdef")}
	in := make([]interface{}, len(expected))
	for i := 0; i < len(expected); i++ {
		in[i] = interface{}(expected[i])
	}
	out := make([][]byte, len(expected))

	query(t, "SELECT ?, ?, ?", in, func(rows *sql.Rows) error { return rows.Scan(&out[0], &out[1], &out[2]) })
	assert.Equal(t, expected, out)
}

func TestBeginTx(t *testing.T) {
	t.Skip("Can not determine current transaction's isolation level: https://bugs.mysql.com/bug.php?id=53341")

	isos := map[sql.IsolationLevel]string{
		sql.LevelReadUncommitted: "READ-UNCOMMITTED",
		sql.LevelReadCommitted:   "READ-COMMITTED",
		sql.LevelRepeatableRead:  "REPEATABLE READ",
		sql.LevelSerializable:    "SERIALIZABLE",
	}

	db := NewDB(t)
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
