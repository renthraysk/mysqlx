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

type queryState int

const (
	queryStart queryState = iota
	queryFetchColumns
	queryFetchedFirstRow
	queryFetchRows
	queryFetchDone
	queryFetchDoneMoreResultSets
	queryError
	queryClosed
)

type rows struct {
	conn    *conn
	state   queryState
	err     error
	names   []string
	columns []*ColumnType
	buf     [16]ColumnType

	t mysqlx.ServerMessages_Type
	b []byte
}

func (r *rows) readColumns(ctx context.Context) error {
	r.state = queryFetchColumns
	r.columns = r.columns[:0]
	r.names = nil

	buf := r.buf[:]
	n := len(buf)

	t, b, err := r.conn.readMessage(ctx)
	for err == nil && t == mysqlx.ServerMessages_RESULTSET_COLUMN_META_DATA {
		if n == 0 {
			n = 16
			buf = make([]ColumnType, n)
		}
		n--
		ct := &buf[n]
		if err := ct.Unmarshal(b); err != nil {
			r.state = queryError
			return err
		}
		r.columns = append(r.columns, ct)
		t, b, err = r.conn.readMessage(ctx)
	}
	if err != nil {
		r.state, r.err = queryError, err
		return err
	}
	switch t {
	case mysqlx.ServerMessages_ERROR:
		r.state, r.err = queryError, newError(b)
		return r.err

	case mysqlx.ServerMessages_RESULTSET_FETCH_DONE:
		r.state = queryFetchDone
		return nil

	case mysqlx.ServerMessages_RESULTSET_FETCH_DONE_MORE_RESULTSETS:
		r.state = queryFetchDoneMoreResultSets
		return nil

	case mysqlx.ServerMessages_RESULTSET_ROW:
		r.state = queryFetchedFirstRow
		r.t, r.b, r.err = t, b, err
		return nil
	}
	return nil
}

func (r *rows) Close() error {
	switch r.state {
	case queryClosed:
	case queryError:
	default:
		t, _, err := r.conn.readMessage(context.Background())
		for err == nil && t != mysqlx.ServerMessages_SQL_STMT_EXECUTE_OK {
			t, _, err = r.conn.readMessage(context.Background())
		}
	}
	r.state = queryClosed
	return nil
}

func (r *rows) Next(values []driver.Value) error {
	t, b, err := r.t, r.b, r.err
	switch r.state {
	case queryFetchedFirstRow:
		r.state = queryFetchRows
	case queryFetchRows:
		t, b, err = r.conn.readMessage(context.Background())
		if err != nil {
			r.state, r.err = queryError, err
			return io.EOF
		}
	default:
		return io.EOF
	}

	switch t {
	case mysqlx.ServerMessages_RESULTSET_FETCH_DONE_MORE_RESULTSETS:
		r.state = queryFetchDoneMoreResultSets
	case mysqlx.ServerMessages_RESULTSET_FETCH_DONE:
		r.state = queryFetchDone
	case mysqlx.ServerMessages_RESULTSET_ROW:
		return r.unmarshalRow(b, values)
	}
	return io.EOF
}

func (r *rows) HasNextResultSet() bool {
	return r.state == queryFetchDoneMoreResultSets
}

func (r *rows) NextResultSet() error {
	if r.state != queryFetchDoneMoreResultSets {
		return io.EOF
	}
	return r.readColumns(context.Background())
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
				t, err := unmarshalTime(b[i:j])
				if err != nil {
					return err
				}
				values[index] = t

			case mysqlx_resultset.ColumnMetaData_BIT:
				bit, nn := binary.Uvarint(b[i:j])
				if nn <= 0 {
					return io.ErrUnexpectedEOF
				}
				values[index] = bit

			default:
				return fmt.Errorf("unknown mysqlx column type %d", column.fieldType)
			}
			i = j
			// Next column
			index++

		default:
			switch tag >> 3 {
			case tagRowValue:
				return fmt.Errorf("Wrong wire type: expected BYTES, got %d", tag&7)
			}

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
