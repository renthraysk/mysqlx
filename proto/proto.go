package proto

import (
	"encoding/binary"
	"math/bits"
)

// Constants that identify the encoding of a value on the wire.
const (
	WireVarint     = 0
	WireFixed32    = 5
	WireFixed64    = 1
	WireBytes      = 2
	WireStartGroup = 3
	WireEndGroup   = 4
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
	return binary.PutUvarint(b, x)
}

// AppendWireString appends a protobuf wire bytes record to p
// tag is assumed to be > 0 and < 16.
func AppendWireString(p []byte, tag uint8, value string) []byte {
	x := uint(len(value))
	n := len(p) + SizeVarint(x)
	if bits.UintSize == 32 {
		p = append(p, tag<<3|WireBytes, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28))
	} else {
		p = append(p, tag<<3|WireBytes, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28)|0x80,
			byte(x>>35)|0x80, byte(x>>42)|0x80, byte(x>>49)|0x80, byte(x>>56), 1)
	}
	p[n] &= 0x7F
	return append(p[:n+1], value...)
}

// AppendWireBytes appends a protobuf wire bytes record to p
// tag is assumed to be > 0 and < 16.
func AppendWireBytes(p []byte, tag uint8, value []byte) []byte {
	x := uint(len(value))
	n := len(p) + SizeVarint(x)
	if bits.UintSize == 32 {
		p = append(p, tag<<3|WireBytes, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28))
	} else {
		p = append(p, tag<<3|WireBytes, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28)|0x80,
			byte(x>>35)|0x80, byte(x>>42)|0x80, byte(x>>49)|0x80, byte(x>>56), 1)
	}
	p[n] &= 0x7F
	return append(p[:n+1], value...)
}
