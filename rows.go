package mysqlx

import (
	"context"
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"time"

	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_resultset"

	"github.com/golang/protobuf/proto"
)

const (
	tagRowValue = 1
)

type Rows interface {
	driver.Rows

	//	HasNextResultSet() bool // driver.RowsNextResultSet
	//	NextResultSet() error

	// ColumnTypeScanType(index int) reflect.Type          // driver.RowsColumnTypeScanType
	ColumnTypeDatabaseTypeName(index int) string        // driver.RowsColumnTypeDatabaseTypeName
	ColumnTypeLength(index int) (length int64, ok bool) // driver.RowsColumnTypeLength
	ColumnTypeNullable(index int) (nullable, ok bool)   // driver.RowsColumnTypeNullable
	//	ColumnTypePrecisionScale(index int) (precision, scale int64, ok bool) // driver.RowsColumnTypePrecisionScale
}

type rows struct {
	conn *conn
	columns
	last struct {
		t   mysqlx.ServerMessages_Type
		b   []byte
		err error
	}

	names []string

	buf [16]ColumnMetaData
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
		r.names = r.columns.Columns()
	}
	return r.names
}

func (r *rows) HasNextResultSet() bool {
	// return r.last.err == nil && r.last.t == mysqlx.ServerMessages_RESULTSET_FETCH_DONE_MORE_RESULTSETS ?
	return false
}

func (r *rows) NextResultSet() error {
	return nil
}

func (r *rows) Next(values []driver.Value) error {

	if r.last.err != nil {
		return r.last.err
	}
	switch r.last.t {

	case mysqlx.ServerMessages_ERROR:
		r.last.err = newError(r.last.b)
		return r.last.err

	case mysqlx.ServerMessages_SQL_STMT_EXECUTE_OK:
		return io.EOF

	case mysqlx.ServerMessages_RESULTSET_FETCH_DONE, mysqlx.ServerMessages_RESULTSET_FETCH_DONE_MORE_RESULTSETS:

	loopUntilExecuteOK:
		r.last.t, r.last.b, r.last.err = r.conn.readMessage(context.TODO())
		switch r.last.t {
		case mysqlx.ServerMessages_NOTICE:
			goto loopUntilExecuteOK

		case mysqlx.ServerMessages_SQL_STMT_EXECUTE_OK:
			return io.EOF
		}

	case mysqlx.ServerMessages_RESULTSET_ROW:
		err := r.unmarshalValues(r.last.b, values)
		r.last.t, r.last.b, r.last.err = r.conn.readMessage(context.TODO())
		return err
	}

	return nil
}

// unmarshalValues parses mysqlx_resultset Row protobuf
func (r *rows) unmarshalValues(b []byte, values []driver.Value) error {

	var s string
	var offset uint64
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
			case mysqlx_resultset.ColumnMetaData_SINT:
				v, nn := binary.Varint(b[i:j])
				if nn <= 0 {
					return io.ErrUnexpectedEOF
				}
				values[index] = v

			case mysqlx_resultset.ColumnMetaData_UINT:
				v, nn := binary.Uvarint(b[i:j])
				if nn <= 0 {
					return io.ErrUnexpectedEOF
				}
				values[index] = v

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

			case mysqlx_resultset.ColumnMetaData_BYTES:
				if column.IsBinary() {
					// database/sql's convertAssign() will cloneBytes() unless being scanned into a sql.RawBytes{}
					values[index] = b[i : j-1 : j-1]
					break
				}
				fallthrough
			case mysqlx_resultset.ColumnMetaData_ENUM:
				// lazy allocation/conversion whole of b to a string... reduces allocations if more > 1 string column.
				if len(s) == 0 {
					s = string(b[i:])
					offset = i
				}
				values[index] = s[i-offset : j-offset-1]

			case mysqlx_resultset.ColumnMetaData_TIME:
				// @TODO

			case mysqlx_resultset.ColumnMetaData_DATETIME:
				v, err := dateTimeToTimeDate(b[i:j])
				if err != nil {
					return err
				}
				values[index] = v

			case mysqlx_resultset.ColumnMetaData_SET:
				// @TODO
				if len(s) == 0 {
					s = string(b[i:])
					offset = i
				}
				values[index] = s[i-offset : j-offset-1]

			case mysqlx_resultset.ColumnMetaData_BIT:
				// @TODO

			case mysqlx_resultset.ColumnMetaData_DECIMAL:
				v, err := decimalToString(b[i:j])
				if err != nil {
					return err
				}
				values[index] = v
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
			}
		}
	}
	if index < len(r.columns) {
		return io.ErrUnexpectedEOF
	}

	return nil
}

func dateTimeToTimeDate(b []byte) (driver.Value, error) {
	year, i := binary.Uvarint(b)
	if i <= 0 {
		return nil, fmt.Errorf("failed to decode datetime year (%#v)", b)
	}
	month, j := binary.Uvarint(b[i:])
	if j <= 0 {
		return nil, fmt.Errorf("failed to decode datetime month (%#v)", b)
	}
	i += j
	day, j := binary.Uvarint(b[i:])
	if j <= 0 {
		return nil, fmt.Errorf("failed to decode datetime day (%#v)", b)
	}
	i += j

	var min, sec, usec uint64

	hour, j := binary.Uvarint(b[i:])
	if j > 0 {
		i += j
		min, j = binary.Uvarint(b[i:])
		if j > 0 {
			i += j
			sec, j = binary.Uvarint(b[i:])
			if j > 0 {
				i += j
				usec, j = binary.Uvarint(b[i:])
			}
		}
	}
	if j < 0 {
		return nil, fmt.Errorf("failed to decode datetime time (%#v)", b)
	}

	return time.Date(int(year), time.Month(month), int(day), int(hour), int(min), int(sec), int(usec)*1000, time.UTC), nil
}

func decimalToString(b []byte) (string, error) {
	if len(b) < 2 {
		return "", fmt.Errorf("failed to parse decimal %#v", b)
	}

	var h uint8

	buf := [96]byte{'-'} // assume negative, easier to slice off if non -'ve
	r := buf[:1]

	for _, l := range b[1:] {
		h = l >> 4
		if h > 9 {
			break
		}
		l &= 0x0F
		if l > 9 {
			r = append(r, '0'+h)
			h = l
			break
		}
		r = append(r, '0'+h, '0'+l)
	}

	// If not negative remove the premptive -
	if h != 0x0B && h != 0x0D {
		r = r[1:]
	}

	if s := b[0]; s > 0 {
		i := len(r) - int(s)
		if i < 0 {
			return "", fmt.Errorf("scale (%d) exceeds precision (%d) in %#v", s, len(r), b)
		}
		r = append(r, 0)
		copy(r[i+1:], r[i:])
		r[i] = '.'
	}
	return string(r), nil
}
