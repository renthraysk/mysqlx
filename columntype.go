package mysqlx

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"

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
	collation      Collation
	hasContentType bool
	contentType    mysqlx_resultset.ContentType_BYTES
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
			collation, nn := binary.Uvarint(b[i:])
			if nn <= 0 {
				return io.ErrUnexpectedEOF
			}
			c.hasCollation = true
			c.collation = Collation(collation)
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
			c.contentType = mysqlx_resultset.ContentType_BYTES(contentType)
			i += uint64(nn)

		default:
			switch tag >> 3 {
			case tagColumnMetaDataName:
				return fmt.Errorf("Wrong wire type: expected BYTES, got %d", tag&7)

			case tagColumnMetaDataType,
				tagColumnMetaDataCollation,
				tagColumnMetaDataFractionalDigits,
				tagColumnMetaDataLength,
				tagColumnMetaDataFlags:
				return fmt.Errorf("Wrong wire type: expected VARINT, got %d", tag&7)
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
				return fmt.Errorf("Unknown wire type (%d)", tag&7)
			}
		}
	}
	if i > n {
		return io.ErrUnexpectedEOF
	}
	return nil
}