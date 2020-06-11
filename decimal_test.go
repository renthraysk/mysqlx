package mysqlx

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

func toBig(x *uint256) *big.Int {
	var y big.Int
	var b [32]byte

	y.SetBytes(x.appendBytes(b[:0]))
	return &y
}

func TestDecimalMulAdd(t *testing.T) {

	tests := []struct {
		x, y, z uint64
	}{
		{0, 10, 1},
		{1, 10, 1},
		{10, 10, 11},
		{math.MaxUint64, math.MaxUint64, 1},
		{math.MaxUint64, math.MaxUint64, math.MaxUint64},
	}

	for _, ts := range tests {

		x := uint256{ts.x}
		x.mulAdd(ts.y, ts.z)

		y := new(big.Int).SetUint64(ts.x)
		y = y.Mul(y, new(big.Int).SetUint64(ts.y))
		y = y.Add(y, new(big.Int).SetUint64(ts.z))

		if a, b := y.String(), toBig(&x).String(); a != b {
			t.Fatalf("failed expected %s, got %s\n", a, b)
		} else {
			fmt.Printf("%s %s\n", a, b)
		}
	}
}
