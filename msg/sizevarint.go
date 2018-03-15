package msg

import "math/bits"

func SizeVarint(x uint64) int {
	// x|1 to prevent Len64() returning 0
	return (bits.Len64(x|1) + 6) / 7
}
