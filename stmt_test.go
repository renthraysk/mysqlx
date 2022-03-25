package mysqlx

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"testing"

	"github.com/renthraysk/mysqlx/msg"
)

func prepare1(cnn *sql.Conn, in any) (out any, err error) {
	stmt, err := cnn.PrepareContext(context.Background(), "SELECT ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(in)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, sql.ErrNoRows
	}
	err = rows.Scan(&out)
	return
}

// value constraint, comparable without string as strings are returned as byte slices.
type value interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~bool
}

func testPrepareValue[T value, S value](t *testing.T, cnn *sql.Conn, in T, expected S) {
	t.Helper()

	t.Run(fmt.Sprintf("%T(%v)", in, in), func(t *testing.T) {
		actual, err := prepare1(cnn, in)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if a, ok := actual.(S); !ok {
			t.Fatalf("expected type %T got %T", expected, actual)
		} else if a != expected {
			t.Fatalf("expected value %T(%v) got %T(%v)", expected, expected, actual, actual)
		}
	})
}

func testPrepareString[T interface{ ~string | ~[]byte }](t *testing.T, cnn *sql.Conn, in T) {

	t.Helper()

	t.Run(fmt.Sprintf("%T(%q)", in, in), func(t *testing.T) {
		actual, err := prepare1(cnn, in)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if a, ok := actual.([]byte); !ok {
			t.Fatalf("expected type []byte got %T", actual)
		} else if string(a) != string(in) {
			t.Fatalf("expected %q got %q", in, actual)
		}
	})
}

func TestPrepare(t *testing.T) {
	connector := NewConnector(t)
	db := sql.OpenDB(connector)
	defer db.Close()

	// Grab single connection
	cnn, err := db.Conn(context.Background())
	assertNoError(t, err)
	defer cnn.Close()

	//	testPrepare(t, cnn, nil)
	testPrepareValue(t, cnn, 0, int64(0))
	testPrepareValue(t, cnn, 1, int64(1))
	testPrepareValue(t, cnn, math.MinInt64, int64(math.MinInt64))
	testPrepareValue(t, cnn, math.MaxInt64, int64(math.MaxInt64))
	testPrepareValue(t, cnn, float64(1.5), float64(1.5))
	testPrepareValue(t, cnn, float32(1.5), float64(1.5))
	testPrepareValue(t, cnn, int64(42), int64(42))

	testPrepareValue(t, cnn, true, int64(1))
	testPrepareValue(t, cnn, false, int64(0))

	testPrepareString(t, cnn, "")
	testPrepareString(t, cnn, "abc")
	testPrepareString(t, cnn, msg.JSONString("{}"))
	testPrepareString(t, cnn, msg.XMLString("<foo />"))

	testPrepareString(t, cnn, []byte{})
	testPrepareString(t, cnn, []byte("abc"))
	testPrepareString(t, cnn, msg.JSON([]byte("{}")))
	testPrepareString(t, cnn, msg.XML([]byte("<foo />")))
}
