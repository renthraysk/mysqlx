package msg

import "math/bits"

// SizeUvarint64 returns the number of bytes required to store a uint64 in base128/varint encoding
func SizeUvarint64(x uint64) int {
	return int(9*uint32(bits.Len64(x))+64) / 64
}

// SizeUvarint32 returns the number of bytes required to store a uint32 in base128/varint encoding
func SizeUvarint32(x uint32) int {
	return int(9*uint32(bits.Len32(x))+64) / 64
}

func SizeUvarint(x uint) int {
	return int(9*uint32(bits.Len(x))+64) / 64
}
