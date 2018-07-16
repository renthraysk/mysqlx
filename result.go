package mysqlx

import (
	"math"

	"github.com/pkg/errors"
)

type result struct {
	hasLastInsertID bool
	lastInsertID    uint64 // protocol defines as uint64, database/sql as int64
	hasRowsAffected bool
	rowsAffected    uint64 // protocol defines as uint64, database/sql as int64
	hasRowsMatched  bool
	rowsMatched     uint64
	hasRowsFound    bool
	rowsFound       uint64
}

// ErrInt64Overflow is the error return when an int64
var ErrInt64Overflow = errors.New("Value exceeded math.MaxInt64")

func (r *result) LastInsertId() (int64, error) {
	if r.lastInsertID > math.MaxInt64 {
		return int64(r.lastInsertID), ErrInt64Overflow
	}
	return int64(r.lastInsertID), nil
}

func (r *result) RowsAffected() (int64, error) {
	if r.rowsAffected > math.MaxInt64 {
		return int64(r.rowsAffected), ErrInt64Overflow
	}
	return int64(r.rowsAffected), nil
}
