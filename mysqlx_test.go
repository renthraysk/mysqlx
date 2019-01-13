package mysqlx

import (
	"context"
	"database/sql"
	"math"
	"testing"
	"time"

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
	// @TODO math.MaxFloat64 appears to get truncated on a roundtrip
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

func TestDuration(t *testing.T) {

	var d []byte

	query(t, "SELECT TIME_FORMAT(?, '%k %i %s')", []interface{}{839*time.Hour - time.Second}, func(rows *sql.Rows) error { return rows.Scan(&d) })

	assert.Equal(t, []byte("838 59 59"), d)
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

func TestRowsAffected(t *testing.T) {
	db := NewDB(t)
	defer db.Close()

	_, err := db.ExecContext(context.Background(), "CREATE TEMPORARY TABLE rowsAffected(ID INT)")
	assert.NoError(t, err)
	{
		r, err := db.ExecContext(context.Background(), "INSERT INTO rowsAffected VALUES(?)", 42)
		assert.NoError(t, err)
		assert.NotNil(t, r)

		n, err := r.RowsAffected()
		assert.NoError(t, err)
		assert.Equal(t, n, int64(1))
	}
	{
		r, err := db.ExecContext(context.Background(), "UPDATE rowsAffected SET ID = ? WHERE ID = ?", 3, 9)
		assert.NoError(t, err)
		assert.NotNil(t, r)

		n, err := r.RowsAffected()
		assert.NoError(t, err)
		assert.Equal(t, n, int64(0))
	}
	{
		r, err := db.ExecContext(context.Background(), "UPDATE rowsAffected SET ID = ? WHERE ID = ?", 3, 42)
		assert.NoError(t, err)
		assert.NotNil(t, r)

		n, err := r.RowsAffected()
		assert.NoError(t, err)
		assert.Equal(t, n, int64(1))
	}
}

func TestMultipleResultsets(t *testing.T) {
	const (
		A int64 = 42
		B       = "testing"
	)
	var (
		a int64
		b string
	)

	db := NewDB(t)
	defer db.Close()

	rows, err := db.Query("CALL spMultipleResultsets(?, ?)", A, B)
	if err != nil {
		t.Fatalf("query failed: %s", err)
	}
	defer rows.Close()

	assert.True(t, rows.Next(), "rows.Next returned false")
	assert.NoError(t, rows.Scan(&a))
	assert.Equal(t, A, a)
	assert.False(t, rows.Next(), "rows.Next returned true")
	assert.True(t, rows.NextResultSet(), "rows.NextResulSet returned false")
	assert.True(t, rows.Next(), "rows.Next returned false")
	assert.NoError(t, rows.Scan(&b))
	assert.Equal(t, B, b)
	assert.False(t, rows.Next(), "rows.Next returned true")
}

func TestBeginTx(t *testing.T) {
	//	t.Skip("Can not determine current transaction's isolation level: https://bugs.mysql.com/bug.php?id=53341")

	isos := []sql.IsolationLevel{
		sql.LevelDefault,
		sql.LevelReadUncommitted,
		sql.LevelReadCommitted,
		sql.LevelRepeatableRead,
		sql.LevelSnapshot,
		sql.LevelSerializable,
	}

	db := NewDB(t)
	defer db.Close()

	for _, level := range isos {
		{
			tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: level})
			assert.NoError(t, err)
			tx.Rollback()
		}
		{
			tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: level, ReadOnly: true})
			assert.NoError(t, err)
			tx.Rollback()
		}
	}
}

func TestQueryTimeout(t *testing.T) {

	db := NewDB(t)
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	_, err := db.ExecContext(ctx, "SELECT SLEEP(3)")
	if err != context.DeadlineExceeded {
		t.Errorf("ExecContext expected to fail with DeadlineExceeded but it returned %v", err)
	}

	{
		var val int64
		rows, err := db.Query("SELECT 42")
		assert.NoError(t, err)
		assert.True(t, rows.Next())
		assert.NoError(t, rows.Scan(&val))
	}
}
