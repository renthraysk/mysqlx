package msg

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"

	"github.com/renthraysk/mysqlx/collation"
	"github.com/renthraysk/mysqlx/proto"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_datatypes"
	"github.com/renthraysk/mysqlx/slice"
)

/*
	Byte banging mysql's X Protocol Any protobuf.
*/

// Tags from Any protobuf
const (
	tagAnyType   = 1
	tagAnyScalar = 2
	tagAnyObject = 3
	tagAnyArray  = 4
)

// Tags from Scalar protobuf
const (
	tagScalarType   = 1
	tagScalarSint   = 2
	tagScalarUint   = 3
	_               // Null
	tagScalarOctets = 5
	tagScalarDouble = 6
	tagScalarFloat  = 7
	tagScalarBool   = 8
	tagScalarString = 9
)

// Tags from Scalar_String protobuf
const (
	tagStringValue     = 1
	tagStringCollation = 2
)

// Tags from Scalar_Octets protobuf
const (
	tagOctetValue       = 1
	tagOctetContentType = 2
)

type AnyAppender interface {
	AppendAny(p []byte, tag uint8) ([]byte, error)
}

// appendAnyUint appends an Any protobuf representing an uint64 value
// tag refers to the protobuf tag index, and is assumed to be > 0 and < 16
func appendAnyUint(p []byte, tag uint8, x uint64) []byte {
	n := proto.SizeVarint64(x)
	p = append(p, tag<<3|proto.WireBytes, 7+byte(n),
		tagAnyType<<3|proto.WireVarint, byte(mysqlx_datatypes.Any_SCALAR),
		tagAnyScalar<<3|proto.WireBytes, 3+byte(n),
		tagScalarType<<3|proto.WireVarint, byte(mysqlx_datatypes.Scalar_V_UINT),
		tagScalarUint<<3|proto.WireVarint, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28)|0x80,
		byte(x>>35)|0x80, byte(x>>42)|0x80, byte(x>>49)|0x80, byte(x>>56), 1)
	n += len(p) - proto.MaxVarintLen64
	p[n-1] &= 0x7F
	return p[:n]
}

// appendAnyInt appends an Any protobuf representing an int64 value
// tag refers to the protobuf tag index, and is assumed to be > 0 and < 16
func appendAnyInt(p []byte, tag uint8, v int64) []byte {
	x := (uint64(v) << 1) ^ uint64(v>>63)
	n := proto.SizeVarint64(x)
	p = append(p, tag<<3|proto.WireBytes, 7+byte(n),
		tagAnyType<<3|proto.WireVarint, byte(mysqlx_datatypes.Any_SCALAR),
		tagAnyScalar<<3|proto.WireBytes, 3+byte(n),
		tagScalarType<<3|proto.WireVarint, byte(mysqlx_datatypes.Scalar_V_SINT),
		tagScalarSint<<3|proto.WireVarint, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28)|0x80,
		byte(x>>35)|0x80, byte(x>>42)|0x80, byte(x>>49)|0x80, byte(x>>56), 1)
	n += len(p) - proto.MaxVarintLen64
	p[n-1] &= 0x7F
	return p[:n]
}

// appendAnyBytes appends an Any protobuf representing an octet ([]byte) value
// tag refers to the protobuf tag index, and is assumed less to be than 16
func appendAnyBytes(p []byte, tag uint8, bytes []byte, contentType contentType) []byte {
	if bytes == nil {
		return appendAnyNull(p, tag)
	}
	n := len(bytes)
	n0 := 1 + proto.SizeVarint(uint(n)) + n // Scalar_Octets size
	if contentType != contentTypePlain {
		n0 += 1 + proto.SizeVarint32(uint32(contentType))
	}
	n1 := 3 + proto.SizeVarint(uint(n0)) + n0 // Scalar size
	n2 := 3 + proto.SizeVarint(uint(n1)) + n1 // Any size

	p, b := slice.ForAppend(p, 1+proto.SizeVarint(uint(n2))+n2)

	i := proto.PutUvarint(b[1:], uint64(n2))
	b[0] = tag<<3 | proto.WireBytes
	b = b[1+i:]
	// Any
	i = proto.PutUvarint(b[3:], uint64(n1))
	b[0] = tagAnyType<<3 | proto.WireVarint
	b[1] = byte(mysqlx_datatypes.Any_SCALAR)
	b[2] = tagAnyScalar<<3 | proto.WireBytes
	b = b[3+i:]
	// Scalar
	i = proto.PutUvarint(b[3:], uint64(n0))
	b[0] = tagScalarType<<3 | proto.WireVarint
	b[1] = byte(mysqlx_datatypes.Scalar_V_OCTETS)
	b[2] = tagScalarOctets<<3 | proto.WireBytes
	b = b[3+i:]

	// Scalar_Octets
	if contentType != contentTypePlain {
		i = proto.PutUvarint(b[1:], uint64(contentType))
		b[0] = tagOctetContentType<<3 | proto.WireVarint
		b = b[1+i:]
	}
	i = proto.PutUvarint(b[1:], uint64(n))
	b[0] = tagOctetValue<<3 | proto.WireBytes
	copy(b[1+i:], bytes)
	return p
}

// appendAnyBytesString appends an Any protobuf representing an octet (string) value
// tag refers to the protobuf tag index, and is assumed less to be than 16
func appendAnyBytesString(p []byte, tag uint8, str string, contentType contentType) []byte {
	n := len(str)
	n0 := 1 + proto.SizeVarint(uint(n)) + n // Scalar_Octets size
	if contentType != contentTypePlain {
		n0 += 1 + proto.SizeVarint32(uint32(contentType))
	}
	n1 := 3 + proto.SizeVarint(uint(n0)) + n0 // Scalar size
	n2 := 3 + proto.SizeVarint(uint(n1)) + n1 // Any size

	p, b := slice.ForAppend(p, 1+proto.SizeVarint(uint(n2))+n2)

	i := proto.PutUvarint(b[1:], uint64(n2))
	b[0] = tag<<3 | proto.WireBytes
	b = b[1+i:]
	// Any
	i = proto.PutUvarint(b[3:], uint64(n1))
	b[0] = tagAnyType<<3 | proto.WireVarint
	b[1] = byte(mysqlx_datatypes.Any_SCALAR)
	b[2] = tagAnyScalar<<3 | proto.WireBytes
	b = b[3+i:]
	// Scalar
	i = proto.PutUvarint(b[3:], uint64(n0))
	b[0] = tagScalarType<<3 | proto.WireVarint
	b[1] = byte(mysqlx_datatypes.Scalar_V_OCTETS)
	b[2] = tagScalarOctets<<3 | proto.WireBytes
	b = b[3+i:]

	// Scalar_Octets
	if contentType != contentTypePlain {
		i = proto.PutUvarint(b[1:], uint64(contentType))
		b[0] = tagOctetContentType<<3 | proto.WireVarint
		b = b[1+i:]
	}
	i = proto.PutUvarint(b[1:], uint64(n))
	b[0] = tagOctetValue<<3 | proto.WireBytes
	copy(b[1+i:], str)
	return p
}

// appendAnyString appends an Any protobuf representing a string value
// tag refers to the protobuf tag index, and is assumed less to be than 16
func appendAnyString(p []byte, tag uint8, s string, collation collation.Collation) []byte {
	n := len(s)
	n0 := 1 + proto.SizeVarint(uint(n)) + n // Scalar_String size
	if collation != 0 {
		n0 += 1 + proto.SizeVarint64(uint64(collation))
	}
	n1 := 3 + proto.SizeVarint(uint(n0)) + n0 // Scalar size
	n2 := 3 + proto.SizeVarint(uint(n1)) + n1 // Any size

	p, b := slice.ForAppend(p, 1+proto.SizeVarint(uint(n2))+n2)

	i := proto.PutUvarint(b[1:], uint64(n2))
	b[0] = tag<<3 | proto.WireBytes
	b = b[1+i:]
	// Any
	i = proto.PutUvarint(b[3:], uint64(n1))
	b[0] = tagAnyType<<3 | proto.WireVarint
	b[1] = byte(mysqlx_datatypes.Any_SCALAR)
	b[2] = tagAnyScalar<<3 | proto.WireBytes
	b = b[3+i:]
	// Scalar
	i = proto.PutUvarint(b[3:], uint64(n0))
	b[0] = tagScalarType<<3 | proto.WireVarint
	b[1] = byte(mysqlx_datatypes.Scalar_V_STRING)
	b[2] = tagScalarString<<3 | proto.WireBytes
	b = b[3+i:]
	// Scalar_String
	if collation != 0 {
		i = proto.PutUvarint(b[1:], uint64(collation))
		b[0] = tagStringCollation<<3 | proto.WireVarint
		b = b[1+i:]
	}
	i = proto.PutUvarint(b[1:], uint64(n))
	b[0] = tagStringValue<<3 | proto.WireBytes
	copy(b[1+i:], s)
	return p
}

// appendAnyFloat64 appends an Any protobuf representing a float64 value
// tag refers to the protobuf tag index, and is assumed to be > 0 and < 16
func appendAnyFloat64(p []byte, tag uint8, f float64) []byte {
	x := math.Float64bits(f)
	return append(p, tag<<3|proto.WireBytes, 15,
		tagAnyType<<3|proto.WireVarint, byte(mysqlx_datatypes.Any_SCALAR),
		tagAnyScalar<<3|proto.WireBytes, 11,
		tagScalarType<<3|proto.WireVarint, byte(mysqlx_datatypes.Scalar_V_DOUBLE),
		tagScalarDouble<<3|proto.WireFixed64, byte(x), byte(x>>8), byte(x>>16), byte(x>>24),
		byte(x>>32), byte(x>>40), byte(x>>48), byte(x>>56))
}

// appendAnyFloat32 appends an Any protobuf representing a float32 value
// tag refers to the protobuf tag index, and is assumed to be > 0 and < 16
func appendAnyFloat32(p []byte, tag uint8, f float32) []byte {
	x := math.Float32bits(f)
	return append(p, tag<<3|proto.WireBytes, 11,
		tagAnyType<<3|proto.WireVarint, byte(mysqlx_datatypes.Any_SCALAR),
		tagAnyScalar<<3|proto.WireBytes, 7,
		tagScalarType<<3|proto.WireVarint, byte(mysqlx_datatypes.Scalar_V_FLOAT),
		tagScalarFloat<<3|proto.WireFixed32, byte(x), byte(x>>8), byte(x>>16), byte(x>>24))
}

// appendAnyBool appends an Any protobuf representing a bool value
// tag refers to the protobuf tag index, and is assumed to be > 0 and < 16
func appendAnyBool(p []byte, tag uint8, b bool) []byte {
	return append(p, tag<<3|proto.WireBytes, 8,
		tagAnyType<<3|proto.WireVarint, byte(mysqlx_datatypes.Any_SCALAR),
		tagAnyScalar<<3|proto.WireBytes, 4,
		tagScalarType<<3|proto.WireVarint, byte(mysqlx_datatypes.Scalar_V_BOOL),
		tagScalarBool<<3|proto.WireVarint, proto.EncodeBool(b))
}

// appendAnyNull appends an Any protobuf representing a NULL/nil value
// tag refers to the protobuf tag index, and is assumed to be > 0 and < 16
func appendAnyNull(p []byte, tag uint8) []byte {
	return append(p, tag<<3|proto.WireBytes, 6,
		tagAnyType<<3|proto.WireVarint, byte(mysqlx_datatypes.Any_SCALAR),
		tagAnyScalar<<3|proto.WireBytes, 2,
		tagScalarType<<3|proto.WireVarint, byte(mysqlx_datatypes.Scalar_V_NULL))
}

func appendAnyTime(p []byte, tag uint8, t time.Time) []byte {
	const fmt = "2006-01-02 15:04:05.000000000"
	const zeroTime = "0000-00-00"

	if t.IsZero() {
		return appendAnyBytesString(p, tag, zeroTime, 0)
	}

	i := len(p)
	p = t.AppendFormat(append(p, tag<<3|proto.WireBytes, 10,
		tagAnyType<<3|proto.WireVarint, byte(mysqlx_datatypes.Any_SCALAR),
		tagAnyScalar<<3|proto.WireBytes, 6,
		tagScalarType<<3|proto.WireVarint, byte(mysqlx_datatypes.Scalar_V_OCTETS),
		tagScalarOctets<<3|proto.WireBytes, 2,
		tagOctetValue<<3|proto.WireBytes, 0), fmt)
	n := len(p) - i - 12
	if n >= 0x80-10 {
		panic("formatted time exceeds 117 bytes in length")
	}
	p[i+11] += byte(n)
	p[i+9] += byte(n)
	p[i+5] += byte(n)
	p[i+1] += byte(n)
	return p
}

const smallsString = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859"

func appendAnyDuration(p []byte, tag uint8, d time.Duration) []byte {
	i := len(p)
	p = append(p, tag<<3|proto.WireBytes, 10,
		tagAnyType<<3|proto.WireVarint, byte(mysqlx_datatypes.Any_SCALAR),
		tagAnyScalar<<3|proto.WireBytes, 6,
		tagScalarType<<3|proto.WireVarint, byte(mysqlx_datatypes.Scalar_V_OCTETS),
		tagScalarOctets<<3|proto.WireBytes, 2,
		tagOctetValue<<3|proto.WireBytes, 0)
	if d < 0 {
		d = -d
		p = append(p, '-')
	}
	p = strconv.AppendUint(p, uint64(d/time.Hour), 10)
	m := 2 * (uint(d/time.Minute) % 60)
	s := 2 * (uint(d/time.Second) % 60)
	p = append(p, ':', smallsString[m], smallsString[m+1], ':', smallsString[s], smallsString[s+1])
	n := len(p) - i - 12
	p[i+11] += byte(n)
	p[i+9] += byte(n)
	p[i+5] += byte(n)
	p[i+1] += byte(n)
	return p
}

func appendAnyValue(p []byte, tag uint8, value interface{}) ([]byte, error) {

derefLoop:
	if value == nil {
		return appendAnyNull(p, tag), nil
	}
	switch v := value.(type) {
	case string:
		return appendAnyString(p, tag, v, 0), nil
	case []byte:
		return appendAnyBytes(p, tag, v, 0), nil
	case uint:
		return appendAnyUint(p, tag, uint64(v)), nil
	case uint8:
		return appendAnyUint(p, tag, uint64(v)), nil
	case uint16:
		return appendAnyUint(p, tag, uint64(v)), nil
	case uint32:
		return appendAnyUint(p, tag, uint64(v)), nil
	case uint64:
		return appendAnyUint(p, tag, v), nil
	case int:
		return appendAnyInt(p, tag, int64(v)), nil
	case int8:
		return appendAnyInt(p, tag, int64(v)), nil
	case int16:
		return appendAnyInt(p, tag, int64(v)), nil
	case int32:
		return appendAnyInt(p, tag, int64(v)), nil
	case int64:
		return appendAnyInt(p, tag, v), nil
	case bool:
		return appendAnyBool(p, tag, v), nil
	case float32:
		return appendAnyFloat32(p, tag, v), nil
	case float64:
		return appendAnyFloat64(p, tag, v), nil
	case time.Time:
		return appendAnyTime(p, tag, v), nil
	case time.Duration:
		return appendAnyDuration(p, tag, v), nil

	default:
		if aa, ok := v.(AnyAppender); ok {
			return aa.AppendAny(p, tag)
		}
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Ptr:
			if rv.IsNil() {
				return appendAnyNull(p, tag), nil
			}
			value = rv.Elem().Interface()
			goto derefLoop

		default:
			return p, fmt.Errorf("unsupported type %T, a %s", value, rv.Kind())
		}
	}
}

func appendAnyValues(p []byte, tag uint8, values []driver.Value) ([]byte, error) {
	var err error

	for i, arg := range values {
		p, err = appendAnyValue(p, tag, arg)
		if err != nil {
			return p, fmt.Errorf("unable to serialize argument %d: %w", i, err)
		}
	}
	return p, nil
}

func appendAnyNamedValues(p []byte, tag uint8, values []driver.NamedValue) ([]byte, error) {
	var err error

	for _, arg := range values {
		if len(arg.Name) > 0 {
			return p, errors.New("mysql does not support the use of named parameters")
		}
		p, err = appendAnyValue(p, tag, arg.Value)
		if err != nil {
			return p, fmt.Errorf("unable to serialize named argument %d: %w", arg.Ordinal, err)
		}
	}
	return p, nil
}
