package msg

import "math/bits"

// SizeUvarint64 returns the number of bytes required to store a uint64 in base128/varint encoding
func SizeUvarint64(x uint64) int {
	// x|1 to prevent Len64() returning 0
	return (bits.Len64(x|1) + 6) / 7
}
