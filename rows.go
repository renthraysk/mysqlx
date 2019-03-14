package mysqlx

import (
	"context"
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"reflect"

	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_resultset"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

const (
	tagRowField = 1
)

type queryState int

const (
	queryStart queryState = iota
	queryFetchColumns
	queryFetchedFirstRow
	queryFetchRows
	queryFetchDone
	queryFetchDoneMoreResultSets
	queryFetchDoneMoreOutParams
	queryError
	queryClosed
)

type rows struct {
	conn      *conn
	state     queryState
	names     []string
	columns   []*ColumnType
	columnBuf [16]ColumnType

	firstRow []byte
}

func (r *rows) readColumns(ctx context.Context) error {
	r.state = queryFetchColumns
	r.columns = r.columns[:0]
	r.names = nil

	buf := r.columnBuf[:]
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
			return errors.Wrap(err, "failed to unmarshal column metadata")
		}
		r.columns = append(r.columns, ct)
		t, b, err = r.conn.readMessage(ctx)
	}
	if err != nil {
		return err
	}
	switch t {
	case mysqlx.ServerMessages_RESULTSET_ROW:
		r.state, r.firstRow = queryFetchedFirstRow, b

	case mysqlx.ServerMessages_ERROR:
		r.state = queryError
		return r.conn.handleError(b)

	case mysqlx.ServerMessages_RESULTSET_FETCH_DONE:
		r.state = queryFetchDone

	case mysqlx.ServerMessages_RESULTSET_FETCH_DONE_MORE_RESULTSETS:
		r.state = queryFetchDoneMoreResultSets
	}
	return nil
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

func (r *rows) Close() error {
	switch r.state {
	case queryClosed, queryError:
	default:
		// We don't know if still holding any values in the buffer, so replace it for closing.
		r.conn.replaceBuffer()
		t, _, err := r.conn.readMessage(context.Background())
		for err == nil && t != mysqlx.ServerMessages_SQL_STMT_EXECUTE_OK {
			t, _, err = r.conn.readMessage(context.Background())
		}
		if err != nil {
			r.state = queryError
			return err
		}
		r.state = queryClosed
	}
	return nil
}

func (r *rows) Next(values []driver.Value) error {
	switch r.state {
	case queryFetchRows:
		t, b, err := r.conn.readMessage(context.Background())
		if err != nil {
			return err
		}
		switch t {
		case mysqlx.ServerMessages_RESULTSET_ROW:
			return r.unmarshalRow(b, values)

		case mysqlx.ServerMessages_RESULTSET_FETCH_DONE:
			r.state = queryFetchDone
		case mysqlx.ServerMessages_RESULTSET_FETCH_DONE_MORE_RESULTSETS:
			r.state = queryFetchDoneMoreResultSets
		case mysqlx.ServerMessages_RESULTSET_FETCH_DONE_MORE_OUT_PARAMS:
			r.state = queryFetchDoneMoreOutParams
		}

	case queryFetchedFirstRow:
		err := r.unmarshalRow(r.firstRow, values)
		r.state = queryFetchRows
		r.firstRow = nil
		return err
	}
	return io.EOF
}

func (r *rows) HasNextResultSet() bool {
	return r.state == queryFetchDoneMoreResultSets || r.state == queryFetchDoneMoreOutParams
}

func (r *rows) NextResultSet() error {

	ctx := context.Background()

	switch r.state {
	case queryFetchDoneMoreResultSets, queryFetchDoneMoreOutParams:
		return r.readColumns(ctx)

	case queryFetchedFirstRow:
		r.state = queryFetchRows
		r.firstRow = nil
		fallthrough
	case queryFetchRows:
		for {
			t, _, err := r.conn.readMessage(ctx)
			if err != nil {
				return err
			}
			switch t {
			case mysqlx.ServerMessages_RESULTSET_ROW:

			case mysqlx.ServerMessages_RESULTSET_FETCH_DONE:
				r.state = queryFetchDone
				return io.EOF

			case mysqlx.ServerMessages_RESULTSET_FETCH_DONE_MORE_RESULTSETS:
				r.state = queryFetchDoneMoreResultSets
				return r.readColumns(ctx)

			case mysqlx.ServerMessages_RESULTSET_FETCH_DONE_MORE_OUT_PARAMS:
				r.state = queryFetchDoneMoreOutParams
				return r.readColumns(ctx)
			}
		}
	}
	return io.EOF
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
		case tagRowField<<3 | proto.WireBytes:
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
			// Value
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
				if column.IsDateTimeDate() {
					var d Date

					if err := d.Unmarshal(b[i:j]); err != nil {
						return err
					}
					values[index] = d
					break
				}
				var dt DateTime
				if err := dt.Unmarshal(b[i:j]); err != nil {
					return err
				}
				values[index] = dt

			case mysqlx_resultset.ColumnMetaData_DECIMAL:
				var d Decimal

				if err := d.Unmarshal(b[i:j]); err != nil {
					return err
				}
				values[index] = d
			case mysqlx_resultset.ColumnMetaData_ENUM:
				values[index] = b[i : j-1 : j-1]

			case mysqlx_resultset.ColumnMetaData_SET:
				values[index] = b[i : j-1 : j-1]

			case mysqlx_resultset.ColumnMetaData_TIME:
				d, err := parseDuration(b[i:j])
				if err != nil {
					return err
				}
				values[index] = d

			case mysqlx_resultset.ColumnMetaData_BIT:
				bit, nn := binary.Uvarint(b[i:j])
				if nn <= 0 {
					return io.ErrUnexpectedEOF
				}
				values[index] = bit

			default:
				return fmt.Errorf("unknown mysqlx column type %s(%d)", column.fieldType.String(), column.fieldType)
			}
			i = j
			// Next column
			index++

		default:
			switch tag >> 3 {
			case tagRowField:
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

func (r *rows) ColumnTypeDatabaseTypeName(index int) string {
	return r.columns[index].DatabaseTypeName()
}

func (r *rows) ColumnTypeLength(index int) (int64, bool) {
	return r.columns[index].Length()
}

func (r *rows) ColumnTypeNullable(index int) (bool, bool) {
	return r.columns[index].Nullable()
}

func (r *rows) ColumnTypePrecisionScale(index int) (int64, int64, bool) {
	return r.columns[index].PrecisionScale()
}

func (r *rows) ColumnTypeScanType(index int) reflect.Type {
	return r.columns[index].ScanType()
}
