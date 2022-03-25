package mysqlx

import (
	"bytes"
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
	"io"
	"math"
	"testing"
	"time"
)

func anySlice[T any](in []T) []any {
	r := make([]any, len(in))
	for i, x := range in {
		r[i] = any(x)
	}
	return r
}

func anySlicePtr[T any](out []T) []any {
	r := make([]any, len(out))
	for i := range out {
		r[i] = &out[i]
	}
	return r
}

var roundTripSelects = [...]string{
	0: "",
	1: "SELECT ?",
	2: "SELECT ?, ?",
	3: "SELECT ?, ?, ?",
	4: "SELECT ?, ?, ?, ?",
	5: "SELECT ?, ?, ?, ?, ?",
	6: "SELECT ?, ?, ?, ?, ?, ?",
	7: "SELECT ?, ?, ?, ?, ?, ?, ?",
	8: "SELECT ?, ?, ?, ?, ?, ?, ?, ?",
	9: "SELECT ?, ?, ?, ?, ?, ?, ?, ?, ?",
}

func roundTrip[T any](t *testing.T, in []T) []T {
	t.Helper()
	out := make([]T, len(in))
	query(t, roundTripSelects[len(in)], anySlice(in), func(rows *sql.Rows) error { return rows.Scan(anySlicePtr(out)...) })
	return out
}

func roundTripComparable[T comparable](t *testing.T, in []T) {
	actual := roundTrip(t, in)

	if len(in) != len(actual) {
		t.Fatalf("slice len expected %d, got %d", len(in), len(actual))
	}
	for i, x := range in {
		if a := actual[i]; a != x {
			t.Fatalf("index %d expected %v, got %v", i, in, a)
		}
	}
}

func TestNull(t *testing.T) {
	out := any(42)
	query(t, roundTripSelects[1], []any{nil}, func(rows *sql.Rows) error { return rows.Scan(&out) })
	if !isNil(out) {
		t.Fatalf("expected nil, got %T(%v)", out, out)
	}
}

func TestTypes(t *testing.T) {
	roundTripComparable(t, []bool{false, true})
	roundTripComparable(t, []uint64{0, math.MaxUint64})
	roundTripComparable(t, []int64{math.MinInt64, 0, math.MaxInt64})
	// @TODO math.MaxFloat32 appears to get truncated on a roundtrip
	roundTripComparable(t, []float32{0, math.SmallestNonzeroFloat32, math.MaxFloat32 - 3.5e+32})
	// @TODO math.MaxFloat64 appears to get truncated on a roundtrip
	roundTripComparable(t, []float64{0, math.SmallestNonzeroFloat64, math.MaxFloat64 - 3.1348623157e+302})
	roundTripComparable(t, []string{"", "abc", "abcdef"})

	//	out := roundTrip(t, [][]byte{{}, {0x00}, []byte("abcdef")})

}

func TestDuration(t *testing.T) {
	var actual []byte

	tests := []struct {
		time.Duration
		expected string
	}{
		{0, "0 00 00"},
		{1 * time.Second, "0 00 01"},
		{59 * time.Second, "0 00 59"},
		{1 * time.Minute, "0 01 00"},
		{59 * time.Minute, "0 59 00"},
		{1 * time.Hour, "1 00 00"},
		{24 * time.Hour, "24 00 00"},
		{839*time.Hour - time.Second, "838 59 59"},

		{-1 * time.Second, "-0 00 01"},
		{-59 * time.Second, "-0 00 59"},
		{-1 * time.Minute, "-0 01 00"},
		{-59 * time.Minute, "-0 59 00"},
		{-1 * time.Hour, "-1 00 00"},
		{-24 * time.Hour, "-24 00 00"},
		{-839*time.Hour + time.Second, "-838 59 59"},
	}

	for _, tt := range tests {
		query(t, "SELECT TIME_FORMAT(?, '%k %i %s')", []any{tt.Duration}, func(rows *sql.Rows) error { return rows.Scan(&actual) })
		if string(actual) != tt.expected {
			t.Fatalf("expected %q got %q", tt.expected, actual)
		}
	}
}

func TestLargeBlob(t *testing.T) {
	const (
		minSize = 1 << 10
		maxSize = 4 << 20
	)

	sizes := []int{minSize, 10240, 1 << 20, maxSize}

	expected := make([]byte, maxSize)

	for _, n := range sizes {
		if _, err := io.ReadFull(rand.Reader, expected[n-minSize:n]); err != nil {
			t.Fatalf("failed to generate random blob: %v", err)
		}
	}

	for _, n := range sizes {
		t.Run(fmt.Sprintf("%d", n), func(t *testing.T) {
			actual := roundTrip(t, [][]byte{expected[:n]})
			if len(actual[0]) != n || !bytes.Equal(expected[:n], actual[0]) {
				t.Fatalf("expected %s...%d got %s...%d", expected[:9], len(expected), actual[0][:9], len(actual[0]))
			}
		})
	}
}

func TestRowsAffected(t *testing.T) {
	db := NewDB(t)
	defer db.Close()

	_, err := db.ExecContext(context.Background(), "DROP TABLE IF EXISTS rowsAffected")
	assertNoError(t, err)

	_, err = db.ExecContext(context.Background(), "CREATE TABLE rowsAffected(ID INT)")
	assertNoError(t, err)
	{
		r, err := db.ExecContext(context.Background(), "INSERT INTO rowsAffected(ID) VALUES(?)", 42)
		assertNoError(t, err)

		n, err := r.RowsAffected()
		assertNoError(t, err)
		if n != int64(1) {
			t.Fatalf("RowsAffected() expected 1, got %v", n)
		}
	}
	{
		r, err := db.ExecContext(context.Background(), "UPDATE rowsAffected SET ID = ? WHERE ID = ?", 3, 9)
		assertNoError(t, err)

		n, err := r.RowsAffected()
		assertNoError(t, err)
		if n != int64(0) {
			t.Fatalf("RowsAffected() expected int64(0), got %T(%v)", n, n)
		}
	}
	{
		r, err := db.ExecContext(context.Background(), "UPDATE rowsAffected SET ID = ? WHERE ID = ?", 3, 42)
		assertNoError(t, err)
		n, err := r.RowsAffected()
		assertNoError(t, err)
		if n != int64(1) {
			t.Fatalf("RowsAffected() expected int64(0), got %T(%v)", n, n)
		}
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
	assertNoError(t, err)
	if !rows.Next() {
		t.Fatal("rows.Next() returned false, expected true")
	}
	assertNoError(t, rows.Scan(&a))
	if a != A {
		t.Fatalf("expected %q got %T(%q)", A, a, a)
	}
	if rows.Next() {
		t.Fatal("rows.Next() returned true, expected false")
	}
	if !rows.NextResultSet() {
		t.Fatal("rows.NextResultSet() returned false, expected true")
	}
	if !rows.Next() {
		t.Fatal("rows.Next() returned false, expected true")
	}
	assertNoError(t, rows.Scan(&b))
	if b != B {
		t.Fatalf("expected %q got %T(%q)", B, b, b)
	}
	if rows.Next() {
		t.Fatal("rows.Next() returned true, expected false")
	}
	assertNoError(t, rows.Close())
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
			assertNoError(t, err)
			tx.Rollback()
		}
		{
			tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: level, ReadOnly: true})
			assertNoError(t, err)
			tx.Rollback()
		}
	}
}

/*
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
		requireNoError(t, err)
		require.True(t, rows.Next())
		requireNoError(t, rows.Scan(&val))
	}
}
*/
