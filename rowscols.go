package mysqlx

import (
	"fmt"
	"reflect"
	"time"

	"github.com/renthraysk/mysqlx/protobuf/mysqlx_resultset"
)

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
	return r.columns[index].length, r.columns[index].hasLength
}

func (r *rows) ColumnTypeNullable(index int) (bool, bool) {
	return r.columns[index].nullable, r.columns[index].hasNullable
}

func (r *rows) ColumnTypePrecisionScale(index int) (int64, int64, bool) {
	c := r.columns[index]
	return c.length, c.scale, c.hasLength && c.hasScale
}

var (
	typeUint    = reflect.TypeOf(uint64(0))
	typeInt     = reflect.TypeOf(int64(0))
	typeBytes   = reflect.TypeOf([]byte{})
	typeFloat32 = reflect.TypeOf(float32(0))
	typeFloat64 = reflect.TypeOf(float64(0))
	typeString  = reflect.TypeOf("")
	typeTime    = reflect.TypeOf(time.Time{})
)

func (r *rows) ColumnTypeScanType(index int) reflect.Type {

	column := r.columns[index]
	switch column.fieldType {
	case mysqlx_resultset.ColumnMetaData_UINT:
		return typeUint

	case mysqlx_resultset.ColumnMetaData_SINT:
		return typeInt

	case mysqlx_resultset.ColumnMetaData_BYTES:
		if column.hasContentType {
			switch column.contentType {
			case mysqlx_resultset.ContentType_BYTES_GEOMETRY:
			case mysqlx_resultset.ContentType_BYTES_JSON:
			case mysqlx_resultset.ContentType_BYTES_XML:
			}
		}
		return typeBytes

	case mysqlx_resultset.ColumnMetaData_DATETIME:
		return typeTime

	case mysqlx_resultset.ColumnMetaData_FLOAT:
		return typeFloat32

	case mysqlx_resultset.ColumnMetaData_DOUBLE:
		return typeFloat64

	case mysqlx_resultset.ColumnMetaData_ENUM:
		return typeString

	case mysqlx_resultset.ColumnMetaData_SET:

	case mysqlx_resultset.ColumnMetaData_BIT:
		return typeUint
	}
	panic(fmt.Sprintf("ColumnTypeScanType: missing support for %s", column.fieldType.String()))
}
