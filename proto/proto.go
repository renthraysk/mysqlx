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

const (
	MaxVarintLen32 = 5
	MaxVarintLen64 = 10
)

// EncodeBool return 1 when b is true, 0 otherwise
func EncodeBool(b bool) byte {
	if b {
		return 1
	}
	return 0
}

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

func AppendWireString(p []byte, tag uint8, value string) []byte {
	x := uint(len(value))
	n := SizeVarint(x)
	if bits.UintSize == 32 {
		p = append(p, tag<<3|WireBytes, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28))
		n += len(p) - MaxVarintLen32
	} else {
		p = append(p, tag<<3|WireBytes, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28)|0x80,
			byte(x>>35)|0x80, byte(x>>42)|0x80, byte(x>>49)|0x80, byte(x>>56)|0x80, 1)
		n += len(p) - MaxVarintLen64
	}
	p[n-1] &= 0x7F
	return append(p[:n], value...)
}

func AppendWireBytes(p []byte, tag uint8, value []byte) []byte {
	x := uint(len(value))
	n := SizeVarint(x)
	if bits.UintSize == 32 {
		p = append(p, tag<<3|WireBytes, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28))
		n += len(p) - MaxVarintLen32
	} else {
		p = append(p, tag<<3|WireBytes, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28)|0x80,
			byte(x>>35)|0x80, byte(x>>42)|0x80, byte(x>>49)|0x80, byte(x>>56)|0x80, 1)
		n += len(p) - MaxVarintLen64
	}
	p[n-1] &= 0x7F
	return append(p[:n], value...)
}
