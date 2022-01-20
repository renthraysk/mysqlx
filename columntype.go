package mysqlx

import (
	"database/sql"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"

	"github.com/renthraysk/mysqlx/collation"
	"github.com/renthraysk/mysqlx/proto"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_resultset"
)

const (
	columnFlagNotNull       = 0x0010
	columnFlagPrimaryKey    = 0x0020
	columnFlagUniqueKey     = 0x0040
	columnFlagMultipleKey   = 0x0080
	columnFlagAutoIncrement = 0x0100
)

const (
	tagColumnMetaDataType             = 1
	tagColumnMetaDataName             = 2
	tagColumnMetaDataCollation        = 8
	tagColumnMetaDataFractionalDigits = 9
	tagColumnMetaDataLength           = 10
	tagColumnMetaDataFlags            = 11
	tagColumnMetaDataContentType      = 12
)

// ColumnType represents metadata of a column
type ColumnType struct {
	fieldType mysqlx_resultset.ColumnMetaData_FieldType
	name      string

	hasNullable bool
	nullable    bool
	hasLength   bool
	length      int64
	hasScale    bool
	scale       int64

	hasCollation   bool
	collation      collation.Collation
	hasContentType bool
	contentType    uint32
}

// Reset resets the metadata for a column, for reusing ColumnType structs
func (c *ColumnType) Reset() {
	c.fieldType = mysqlx_resultset.ColumnMetaData_SINT

	c.name = ""
	c.hasNullable = false
	c.hasLength = false
	c.hasScale = false
	c.hasCollation = false
	c.hasContentType = false
}

func (c *ColumnType) DatabaseTypeName() string {
	return c.fieldType.String()
}

func (c *ColumnType) Length() (int64, bool) {
	return c.length, c.hasLength
}

func (c *ColumnType) Nullable() (bool, bool) {
	return c.nullable, c.hasNullable
}

func (c *ColumnType) PrecisionScale() (int64, int64, bool) {
	return c.length, c.scale, c.hasLength && c.hasScale
}

type NullUint64 struct {
	Value uint64
	Valid bool
}

func (n *NullUint64) Scan(src any) error {
	if src == nil {
		n.Valid = false
		return nil
	}
	switch v := src.(type) {
	case uint64:
		n.Value = v
	case uint32:
		n.Value = uint64(v)
	case uint16:
		n.Value = uint64(v)
	case uint8:
		n.Value = uint64(v)
	case uint:
		n.Value = uint64(v)
	case []byte:
		u, i := binary.Uvarint(v)
		n.Valid = i > 0
		n.Value = u
	default:
		return fmt.Errorf("unable to convert type %T to NullInt64", src)
	}
	return nil
}

var (
	typeUint         = reflect.TypeOf(uint64(0))
	typeNullUint64   = reflect.TypeOf(NullUint64{})
	typeInt          = reflect.TypeOf(int64(0))
	typeNullInt64    = reflect.TypeOf(sql.NullInt64{})
	typeBytes        = reflect.TypeOf([]byte{})
	typeFloat32      = reflect.TypeOf(float32(0))
	typeFloat64      = reflect.TypeOf(float64(0))
	typeNullFloat64  = reflect.TypeOf(sql.NullFloat64{})
	typeString       = reflect.TypeOf("")
	typeNullString   = reflect.TypeOf(sql.NullString{})
	typeDuration     = typeInt
	typeNullDuration = typeNullInt64
	typeDate         = reflect.TypeOf(Date{})
	typeNullDate     = reflect.TypeOf(NullDate{})
	typeDateTime     = reflect.TypeOf(DateTime{})
	typeNullDateTime = reflect.TypeOf(NullDateTime{})
	typeAny          = reflect.TypeOf(new(any)).Elem()
)

func (c *ColumnType) ScanType() reflect.Type {

	if c.hasNullable && c.nullable {
		switch c.fieldType {
		case mysqlx_resultset.ColumnMetaData_UINT:
			return typeNullUint64
		case mysqlx_resultset.ColumnMetaData_SINT:
			return typeNullInt64

		case mysqlx_resultset.ColumnMetaData_BYTES:
			if c.hasContentType {
				switch mysqlx_resultset.ContentType_BYTES(c.contentType) {
				case mysqlx_resultset.ContentType_BYTES_GEOMETRY:
				case mysqlx_resultset.ContentType_BYTES_JSON:
				case mysqlx_resultset.ContentType_BYTES_XML:
				}
			}
			// @TODO ?!
			if c.hasCollation && !c.collation.IsBinary() {
				return typeNullString
			}
			return typeBytes

		case mysqlx_resultset.ColumnMetaData_DATETIME:
			if c.hasContentType &&
				mysqlx_resultset.ContentType_DATETIME(c.contentType) == mysqlx_resultset.ContentType_DATETIME_DATE {
				return typeNullDate
			}
			return typeNullDateTime

		case mysqlx_resultset.ColumnMetaData_TIME:
			return typeNullDuration

		case mysqlx_resultset.ColumnMetaData_FLOAT:
			return typeNullFloat64 // @TODO 32?

		case mysqlx_resultset.ColumnMetaData_DOUBLE:
			return typeNullFloat64

		case mysqlx_resultset.ColumnMetaData_ENUM:
			return typeNullString
		}
		return typeAny
	}

	switch c.fieldType {
	case mysqlx_resultset.ColumnMetaData_UINT:
		return typeUint

	case mysqlx_resultset.ColumnMetaData_SINT:
		return typeInt

	case mysqlx_resultset.ColumnMetaData_BYTES:
		if c.hasContentType {
			switch mysqlx_resultset.ContentType_BYTES(c.contentType) {
			case mysqlx_resultset.ContentType_BYTES_GEOMETRY:
			case mysqlx_resultset.ContentType_BYTES_JSON:
			case mysqlx_resultset.ContentType_BYTES_XML:
			}
		}
		// @TODO ?!
		if c.hasCollation && !c.collation.IsBinary() {
			return typeString
		}
		return typeBytes

	case mysqlx_resultset.ColumnMetaData_DATETIME:
		if c.hasContentType &&
			mysqlx_resultset.ContentType_DATETIME(c.contentType) == mysqlx_resultset.ContentType_DATETIME_DATE {
			return typeDate
		}
		return typeDateTime

	case mysqlx_resultset.ColumnMetaData_TIME:
		return typeDuration

	case mysqlx_resultset.ColumnMetaData_FLOAT:
		return typeFloat32

	case mysqlx_resultset.ColumnMetaData_DOUBLE:
		return typeFloat64

	case mysqlx_resultset.ColumnMetaData_ENUM:
		return typeString

	case mysqlx_resultset.ColumnMetaData_SET:
		return typeBytes

	case mysqlx_resultset.ColumnMetaData_BIT:
		return typeUint
	}
	return typeAny
}

// IsBinary returns true if column represets a binary type
func (c *ColumnType) IsBinary() bool {
	return c.fieldType == mysqlx_resultset.ColumnMetaData_BYTES && c.hasCollation && c.collation.IsBinary()
}

// Unmarshal unmarhals the mysqlx_resultset.ColumnMetaData protobuf into the ColumnType
func (c *ColumnType) Unmarshal(b []byte) error {
	var nn int

	c.Reset()

	i, n := uint64(0), uint64(len(b))
	for i < n {
		tag := uint64(b[i])
		i++
		switch tag {
		case tagColumnMetaDataType<<3 | proto.WireVarint:
			fieldType, nn := binary.Uvarint(b[i:])
			if nn <= 0 {
				return io.ErrUnexpectedEOF
			}
			c.fieldType = mysqlx_resultset.ColumnMetaData_FieldType(fieldType)
			i += uint64(nn)

		case tagColumnMetaDataName<<3 | proto.WireBytes:
			j, nn := binary.Uvarint(b[i:])
			if nn <= 0 {
				return io.ErrUnexpectedEOF
			}
			i += uint64(nn)
			j += i
			if j > n {
				return io.ErrUnexpectedEOF
			}
			c.name = string(b[i:j])
			i = j

		case tagColumnMetaDataCollation<<3 | proto.WireVarint:
			col, nn := binary.Uvarint(b[i:])
			if nn <= 0 {
				return io.ErrUnexpectedEOF
			}
			c.hasCollation = true
			c.collation = collation.Collation(col)
			i += uint64(nn)

		case tagColumnMetaDataFractionalDigits<<3 | proto.WireVarint:
			s, nn := binary.Uvarint(b[i:])
			if nn <= 0 {
				return io.ErrUnexpectedEOF
			}
			c.hasScale = true
			c.scale = int64(s)
			i += uint64(nn)

		case tagColumnMetaDataLength<<3 | proto.WireVarint:
			length, nn := binary.Uvarint(b[i:])
			if nn <= 0 {
				return io.ErrUnexpectedEOF
			}
			if length > math.MaxInt64 {
				return errors.New("Length exceeds MaxInt64")
			}
			c.hasLength = true
			c.length = int64(length)
			i += uint64(nn)

		case tagColumnMetaDataFlags<<3 | proto.WireVarint:
			flags, nn := binary.Uvarint(b[i:])
			if nn <= 0 {
				return io.ErrUnexpectedEOF
			}
			c.hasNullable = true
			c.nullable = flags&columnFlagNotNull == 0
			i += uint64(nn)

		case tagColumnMetaDataContentType<<3 | proto.WireVarint:
			contentType, nn := binary.Uvarint(b[i:])
			if nn <= 0 {
				return io.ErrUnexpectedEOF
			}
			c.hasContentType = true
			c.contentType = uint32(contentType)
			i += uint64(nn)

		default:
			switch tag >> 3 {
			case tagColumnMetaDataName:
				return fmt.Errorf("wrong wire type: expected BYTES, got %d", tag&7)

			case tagColumnMetaDataType,
				tagColumnMetaDataCollation,
				tagColumnMetaDataFractionalDigits,
				tagColumnMetaDataLength,
				tagColumnMetaDataFlags:
				return fmt.Errorf("wrong wire type: expected VARINT, got %d", tag&7)
			}

			if tag > 0x7F {
				i--
				tag, nn = binary.Uvarint(b[i:])
				if nn <= 0 {
					return io.ErrUnexpectedEOF
				}
				i += uint64(nn)
			}

			switch tag & 7 {
			case proto.WireVarint:
				_, nn = binary.Uvarint(b[i:])
				if nn <= 0 {
					return io.ErrUnexpectedEOF
				}
				i += uint64(nn)
			case proto.WireFixed64:
				i += 8
			case proto.WireBytes:
				l, nn := binary.Uvarint(b[i:])
				if nn <= 0 {
					return io.ErrUnexpectedEOF
				}
				i += uint64(nn)
				i += l
			case proto.WireFixed32:
				i += 4
			default:
				return fmt.Errorf("unknown wire type (%d)", tag&7)
			}
		}
	}
	if i > n {
		return io.ErrUnexpectedEOF
	}
	return nil
}
