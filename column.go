package mysqlx

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/golang/protobuf/proto"

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

type columnMetaData struct {
	fieldType mysqlx_resultset.ColumnMetaData_FieldType

	name              string
	hasFlags          bool
	flags             uint32
	hasLength         bool
	length            uint32
	hasCollation      bool
	collation         Collation
	hasPrecisionScale bool
	precision         int64
	scale             int64
	hasContentType    bool
	contentType       uint32
}

func (c *columnMetaData) reset() {
	c.fieldType = mysqlx_resultset.ColumnMetaData_SINT
	c.name = ""
	c.hasFlags = false
	c.hasLength = false
	c.hasPrecisionScale = false
	c.hasCollation = false
	c.hasContentType = false
}

func (c *columnMetaData) isBinary() bool {
	return c.fieldType == mysqlx_resultset.ColumnMetaData_BYTES && c.hasCollation && c.collation.IsBinary()
}

func (c *columnMetaData) nullable() (bool, bool) {
	return c.flags&columnFlagNotNull == 0, c.hasFlags
}

func (c *columnMetaData) unmarshal(b []byte) error {
	var nn int

	c.reset()

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
			_, nn = binary.Uvarint(b[i:])

			// @TODO
			i += uint64(nn)
		case tagColumnMetaDataLength<<3 | proto.WireVarint:
			length, nn := binary.Uvarint(b[i:])
			if nn <= 0 {
				return io.ErrUnexpectedEOF
			}
			c.hasLength = true
			c.length = uint32(length)
			i += uint64(nn)
		case tagColumnMetaDataFlags<<3 | proto.WireVarint:
			flags, nn := binary.Uvarint(b[i:])
			if nn <= 0 {
				return io.ErrUnexpectedEOF
			}
			c.hasFlags = true
			c.flags = uint32(flags)
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

	return nil
}
