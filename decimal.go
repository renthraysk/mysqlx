package mysqlx

import (
	"encoding/binary"
	"math/bits"
)

const (
	ErrDecimalTooShort = errorString("decimal too short")
)

type uint256 [4]uint64

// appendBytes returns the number in big endian order appended to buf
func (x *uint256) appendBytes(buf []byte) []byte {
	var b [32]byte

	binary.BigEndian.PutUint64(b[0:], x[3])
	binary.BigEndian.PutUint64(b[8:], x[2])
	binary.BigEndian.PutUint64(b[16:], x[1])
	binary.BigEndian.PutUint64(b[24:], x[0])
	i := 0
	for i < 31 && b[i] == 0 {
		i++
	}
	return append(buf, b[i:]...)
}

// mulAdd x = x*y + z returns non zero if overflowed
func (x *uint256) mulAdd(y uint64, z uint64) uint64 {
	var h0, h1, h2, h3, c uint64

	h0, x[0] = bits.Mul64(x[0], y)
	h1, x[1] = bits.Mul64(x[1], y)
	h2, x[2] = bits.Mul64(x[2], y)
	h3, x[3] = bits.Mul64(x[3], y)
	x[0], c = bits.Add64(x[0], z, 0)
	x[1], c = bits.Add64(x[1], h0, c)
	x[2], c = bits.Add64(x[2], h1, c)
	x[3], c = bits.Add64(x[3], h2, c)
	return h3 | c
}

const (
	maxLength = 77 // Maximum number of digits in a MySQL DECIMAL
	digits    = "0123456789"
)

type Decimal interface {
	Decompose(buf []byte) (form byte, negative bool, coefficient []byte, exponent int32)
	String() string
}

type decimal []byte

func (d decimal) Decompose(buf []byte) (form byte, negative bool, coefficient []byte, exponent int32) {
	var ui256 uint256

	form = 0
	negative = false
	exponent = -int32(d[0])
	for _, x := range d[1:] {
		y := x & 0x0F
		x >>= 4
		if x > 9 {
			negative = x == 0xB || x == 0xD
			break
		}
		ui256.mulAdd(10, uint64(x))
		if y > 9 {
			negative = y == 0xB || y == 0xD
			break
		}
		ui256.mulAdd(10, uint64(y))
	}
	coefficient = ui256.appendBytes(buf[:0])
	return
}

func (d decimal) String() string {

	buf := [1 + maxLength + 1]byte{0: '-'}
	b := buf[:1]
	for _, x := range d[1:] {
		y := x & 0x0F
		x >>= 4
		if x > 9 {
			if x != 0xB && x != 0xD {
				b = b[1:]
			}
			break
		}
		if y > 9 {
			b = append(b, digits[x])
			if y != 0xB && y != 0xD {
				b = b[1:]
			}
			break
		}
		b = append(b, digits[x], digits[y])
	}

	// Scale
	if s := int(d[0]); s > 0 {
		// @TODO error?
		if i := len(b) - s; i >= 0 {
			b = append(b, 0)
			copy(b[i+1:], b[i:])
			b[i] = '.'
		}
	}
	return string(b)
}
