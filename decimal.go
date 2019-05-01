package mysqlx

import (
	"encoding/binary"
	"math/bits"
)

const (
	ErrDecimalTooShort = errorString("decimal too short")
)

type Uint256 [4]uint64

// Add x = x+y, returning carry
func (x *Uint256) Add(y *Uint256) uint64 {
	var c uint64
	x[0], c = bits.Add64(x[0], y[0], 0)
	x[1], c = bits.Add64(x[1], y[1], c)
	x[2], c = bits.Add64(x[2], y[2], c)
	x[3], c = bits.Add64(x[3], y[3], c)
	return c
}

// AddUint x = x+y, returning carry
func (x *Uint256) AddUint(y uint64) uint64 {
	var c uint64
	x[0], c = bits.Add64(x[0], y, 0)
	x[1], c = bits.Add64(x[1], 0, c)
	x[2], c = bits.Add64(x[2], 0, c)
	x[3], c = bits.Add64(x[3], 0, c)
	return c
}

// AppendBytes returns the number in little endian order appended to buf
func (x *Uint256) AppendBytes(buf []byte) []byte {
	var b [32]byte

	binary.LittleEndian.PutUint64(b[0:], x[0])
	binary.LittleEndian.PutUint64(b[8:], x[1])
	binary.LittleEndian.PutUint64(b[16:], x[2])
	binary.LittleEndian.PutUint64(b[24:], x[3])
	i := 31
	for i > 0 && b[i] == 0 {
		i--
	}
	i++
	return append(buf, b[:i]...)
}

// mul10add x = x*10 + y
func mul10Add(x *Uint256, y uint64) uint64 {
	// @TODO Overflow checking
	x.Add(x) // *2
	t := *x
	t.Add(&t) // *4
	t.Add(&t) // *8
	x.Add(&t) // 2 + 8
	return x.AddUint(y)
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
	var ui256 Uint256

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
		mul10Add(&ui256, uint64(x))
		if y > 9 {
			negative = y == 0xB || y == 0xD
			break
		}
		mul10Add(&ui256, uint64(y))
	}
	coefficient = ui256.AppendBytes(buf[:0])
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
