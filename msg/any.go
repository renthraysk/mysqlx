package msg

import (
	"math"

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

// appendAnyUint appends an Any protobuf representing an uint64 value
// tag refers to the protobuf tag index, and is assumed to be > 0 and < 16
func appendAnyUint(p []byte, tag uint8, x uint64) []byte {
	n := proto.SizeVarint64(x)
	p = append(p, tag<<3|proto.WireBytes, 7+byte(n),
		tagAnyType<<3|proto.WireVarint, byte(mysqlx_datatypes.Any_SCALAR),
		tagAnyScalar<<3|proto.WireBytes, 3+byte(n),
		tagScalarType<<3|proto.WireVarint, byte(mysqlx_datatypes.Scalar_V_UINT),
		tagScalarUint<<3|proto.WireVarint, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28)|0x80,
		byte(x>>35)|0x80, byte(x>>42)|0x80, byte(x>>49)|0x80, byte(x>>56)|0x80, 1)
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
		byte(x>>35)|0x80, byte(x>>42)|0x80, byte(x>>49)|0x80, byte(x>>56)|0x80, 1)
	n += len(p) - proto.MaxVarintLen64
	p[n-1] &= 0x7F
	return p[:n]
}

// appendAnyBytes appends an Any protobuf representing an octet ([]byte) value
// tag refers to the protobuf tag index, and is assumed less to be than 16
func appendAnyBytes(p []byte, tag uint8, bytes []byte, contentType ContentType) []byte {
	if bytes == nil {
		return appendAnyNull(p, tag)
	}
	n := len(bytes)
	n0 := 1 + proto.SizeVarint(uint(n)) + n // Scalar_Octets size
	if contentType != ContentTypePlain {
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
	if contentType != ContentTypePlain {
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
func appendAnyBytesString(p []byte, tag uint8, str string, contentType ContentType) []byte {
	n := len(str)
	n0 := 1 + proto.SizeVarint(uint(n)) + n // Scalar_Octets size
	if contentType != ContentTypePlain {
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
	if contentType != ContentTypePlain {
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
