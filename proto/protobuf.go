package proto

import (
	"encoding/binary"
	"math/bits"
)

// Constants that identify the encoding of a value on the wire.
const (
	WireVarint     = 0
	WireFixed64    = 1
	WireBytes      = 2
	WireStartGroup = 3
	WireEndGroup   = 4
	WireFixed32    = 5
)

// SizeVarint64 returns the number of bytes required to store a uint64 in base128/varint encoding
func SizeVarint64(x uint64) int {
	return int(9*uint32(bits.Len64(x))+64) / 64
}

// Sizevarint32 returns the number of bytes required to store a uint32 in base128/varint encoding
func SizeVarint32(x uint32) int {
	return int(9*uint32(bits.Len32(x))+64) / 64
}

// SizeVarint returns the number of bytes required to store a uint in base128/varint encoding
func SizeVarint(x uint) int {
	return int(9*uint32(bits.Len(x))+64) / 64
}

func PutUvarint(b []byte, x uint64) int {
	if x < 0x80 {
		b[0] = byte(x)
		return 1
	}
	return binary.PutUvarint(b, x)
}
