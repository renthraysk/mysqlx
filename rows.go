package mysqlx

import (
	"context"
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"io"
	"math"

	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_resultset"

	"github.com/golang/protobuf/proto"
)

const (
	tagRowValue = 1
)

type rows struct {
	conn    *conn
	columns []*columnMetaData
	last    struct {
		t   mysqlx.ServerMessages_Type
		b   []byte
		err error
	}

	names []string

	buf [16]columnMetaData
}

func (r *rows) Close() error {
	for r.last.err == nil && r.last.t != mysqlx.ServerMessages_SQL_STMT_EXECUTE_OK {
		r.last.t, r.last.b, r.last.err = r.conn.readMessage(context.TODO())
	}
	r.conn = nil
	return r.last.err
}

func (r *rows) Columns() []string {
	if r.names == nil {
		r.names = make([]string, len(r.columns))
		for index, column := range r.columns {
			r.names[index] = column.name
		}
	}
	return r.names
}

func (r *rows) ColumnTypeDatabaseTypeName(index int) string {
	return r.columns[index].fieldType.String()
}

func (r *rows) ColumnTypeLength(index int) (int64, bool) {
	return int64(r.columns[index].length), r.columns[index].hasLength
}

func (r *rows) ColumnTypeNullable(index int) (bool, bool) {
	return r.columns[index].nullable()
}

func (r *rows) Next(values []driver.Value) error {

	if r.last.err != nil {
		return r.last.err
	}
	switch r.last.t {
	case mysqlx.ServerMessages_RESULTSET_ROW:
		err := r.unmarshalRow(r.last.b, values)
		r.last.t, r.last.b, r.last.err = r.conn.readMessage(context.TODO())
		return err

	case mysqlx.ServerMessages_ERROR:
		r.last.err = newError(r.last.b)
		return r.last.err

	case mysqlx.ServerMessages_SQL_STMT_EXECUTE_OK:
		r.last.err = io.EOF
		return r.last.err

	case mysqlx.ServerMessages_RESULTSET_FETCH_DONE, mysqlx.ServerMessages_RESULTSET_FETCH_DONE_MORE_RESULTSETS:

	loopUntilExecuteOK:
		r.last.t, r.last.b, r.last.err = r.conn.readMessage(context.TODO())
		switch r.last.t {
		case mysqlx.ServerMessages_NOTICE:
			goto loopUntilExecuteOK

		case mysqlx.ServerMessages_SQL_STMT_EXECUTE_OK:
			return io.EOF
		}

	}

	return nil
}

// unmarshalRow parses mysqlx_resultset Row protobuf
func (r *rows) unmarshalRow(b []byte, values []driver.Value) error {
	var j uint64
	var nn int

	i := uint64(0)
	n := uint64(len(b))

	// Column index
	index := 0

	// Breaks as soon as parsed a value per column even if hasn't parsed entire protobuf
	for i < n && index < len(r.columns) {
		tag := uint64(b[i])
		i++
		if i >= n {
			return io.ErrUnexpectedEOF
		}
		switch tag {
		case tagRowValue<<3 | proto.WireBytes:
			// Length...
			j = uint64(b[i])
			i++
			if j > 0x7F {
				i--
				j, nn = binary.Uvarint(b[i:])
				if nn <= 0 {
					return io.ErrUnexpectedEOF
				}
				i += uint64(nn)
			}
			// Length == 0 means nil
			if j == 0 {
				values[index] = nil
				index++
				continue
			}
			j += i
			if j > n {
				return io.ErrUnexpectedEOF
			}
			switch column := r.columns[index]; column.fieldType {
			case mysqlx_resultset.ColumnMetaData_UINT:
				v, nn := binary.Uvarint(b[i:j])
				if nn <= 0 {
					return io.ErrUnexpectedEOF
				}
				values[index] = v

			case mysqlx_resultset.ColumnMetaData_SINT:
				v, nn := binary.Varint(b[i:j])
				if nn <= 0 {
					return io.ErrUnexpectedEOF
				}
				values[index] = v

			case mysqlx_resultset.ColumnMetaData_BYTES:
				/*
					if column.hasContentType {
						// contentType is defined as a uint32 in mysqlx_resultset.proto
						// But the enum of possible values is assigned type int32 by protoc
						switch column.contentType {
						case uint32(mysqlx_resultset.ContentType_BYTES_GEOMETRY):
							values[index] = b[i : j-1 : j-1]

						case uint32(mysqlx_resultset.ContentType_BYTES_JSON):
							values[index] = b[i : j-1 : j-1]

						case uint32(mysqlx_resultset.ContentType_BYTES_XML):
							values[index] = b[i : j-1 : j-1]

						default:
							values[index] = b[i : j-1 : j-1]
						}
						break
					}
				*/
				values[index] = b[i : j-1 : j-1]

			case mysqlx_resultset.ColumnMetaData_DOUBLE:
				if j-i != 8 {
					return io.ErrUnexpectedEOF
				}
				values[index] = math.Float64frombits(binary.LittleEndian.Uint64(b[i:j]))

			case mysqlx_resultset.ColumnMetaData_FLOAT:
				if j-i != 4 {
					return io.ErrUnexpectedEOF
				}
				values[index] = math.Float32frombits(binary.LittleEndian.Uint32(b[i:j]))

			case mysqlx_resultset.ColumnMetaData_DATETIME:
				/*
					if column.hasContentType {
						switch column.contentType {
						case uint32(mysqlx_resultset.ContentType_DATETIME_DATE):
						case uint32(mysqlx_resultset.ContentType_DATETIME_DATETIME):
							// Not defined in protobuf
							const MysqlxColumnFlagsDateTimeTimeStamp = 1
							if column.hasFlags && column.flags&MysqlxColumnFlagsDateTimeTimeStamp != 0 {
							}
						}
					}*/
				t, err := unmarshalDateTime(b[i:j])
				if err != nil {
					return err
				}
				values[index] = t

			case mysqlx_resultset.ColumnMetaData_DECIMAL:
				d, err := unmarshalDecimal(b[i:j])
				if err != nil {
					return err
				}
				values[index] = d

			case mysqlx_resultset.ColumnMetaData_ENUM:
				values[index] = b[i : j-1 : j-1]

			case mysqlx_resultset.ColumnMetaData_SET:
				values[index] = b[i : j-1 : j-1]

			case mysqlx_resultset.ColumnMetaData_TIME:
				t, err := unmarshalTime(b[i : j-1])
				if err != nil {
					return err
				}
				values[index] = t

			case mysqlx_resultset.ColumnMetaData_BIT:
				bit, err := unmarshalBit(b[i : j-1])
				if err != nil {
					return err
				}
				values[index] = bit

			default:
				return fmt.Errorf("unknown mysqlx column type %d", column.fieldType)
			}
			i = j
			// Next column
			index++

		default:
			// Skip over tags & values not familar with
			if tag > 0x7F {
				i--
				tag, nn = binary.Uvarint(b[i:])
				if nn <= 0 {
					return io.ErrUnexpectedEOF
				}
				i += uint64(nn)
			}

			switch tag >> 3 {
			case tagRowValue:
				return fmt.Errorf("Wrong wire type: expected BYTES, got %d", tag&7)
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
				j, nn = binary.Uvarint(b[i:])
				if nn <= 0 {
					return io.ErrUnexpectedEOF
				}
				i += uint64(nn)
				i += j
			case proto.WireFixed32:
				i += 4
			default:
				return fmt.Errorf("Unknown wire type (%d)", tag&7)
			}
		}
	}
	if index < len(r.columns) {
		return io.ErrUnexpectedEOF
	}

	return nil
}
