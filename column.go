package mysqlx

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"reflect"

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

type ColumnType struct {
	fieldType        mysqlx_resultset.ColumnMetaData_FieldType
	Name             string
	DatabaseTypeName string

	HasNullable       bool
	Nullable          bool
	HasLength         bool
	Length            int64
	HasPrecisionScale bool
	Precision         int64
	Scale             int64
	ScanType          reflect.Type

	hasCollation   bool
	collation      Collation
	hasContentType bool
	contentType    uint32
}

func (c *ColumnType) Reset() {
	c.fieldType = mysqlx_resultset.ColumnMetaData_SINT
	c.DatabaseTypeName = c.fieldType.String()
	c.Name = ""
	c.HasNullable = false
	c.HasLength = false
	c.HasPrecisionScale = false
	c.hasCollation = false
	c.hasContentType = false
}

func (c *ColumnType) IsBinary() bool {
	return c.fieldType == mysqlx_resultset.ColumnMetaData_BYTES && c.hasCollation && c.collation.IsBinary()
}

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
			c.DatabaseTypeName = c.fieldType.String()
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
			c.Name = string(b[i:j])
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
			c.HasPrecisionScale = true
			c.Scale = int64(s)
			i += uint64(nn)

		case tagColumnMetaDataLength<<3 | proto.WireVarint:
			length, nn := binary.Uvarint(b[i:])
			if nn <= 0 {
				return io.ErrUnexpectedEOF
			}
			if length > math.MaxInt64 {
				return errors.New("Length exceeds MaxInt64")
			}
			c.HasLength = true
			c.Length = int64(length)
			i += uint64(nn)

		case tagColumnMetaDataFlags<<3 | proto.WireVarint:
			flags, nn := binary.Uvarint(b[i:])
			if nn <= 0 {
				return io.ErrUnexpectedEOF
			}
			c.HasNullable = true
			c.Nullable = flags&columnFlagNotNull == 0
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
			if tag > 0x7F {
				i--
				tag, nn = binary.Uvarint(b[i:])
				if nn <= 0 {
					return io.ErrUnexpectedEOF
				}
				i += uint64(nn)
			}

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
	c.HasPrecisionScale = c.HasPrecisionScale && c.HasLength
	c.Precision = c.Length
	return nil
}
