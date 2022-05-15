package msg

import (
	"github.com/renthraysk/mysqlx/collation"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_resultset"
)

type contentType uint32

const (
	contentTypePlain    contentType = 0
	contentTypeGeometry             = contentType(mysqlx_resultset.ContentType_BYTES_GEOMETRY)
	contentTypeJSON                 = contentType(mysqlx_resultset.ContentType_BYTES_JSON)
	contentTypeXML                  = contentType(mysqlx_resultset.ContentType_BYTES_XML)
)

type String struct {
	Value     string
	Collation collation.Collation
}

func (s *String) AppendAny(p []byte, tag uint8) ([]byte, error) {
	return appendAnyString(p, tag, s.Value, s.Collation), nil
}

func Null() null { return null{} }

type null struct{}

func (n null) AppendAny(p []byte, tag uint8) ([]byte, error) {
	return appendAnyNull(p, tag), nil
}

type geometry[T []byte | string] struct{ value T }

func Geometry[T []byte | string](value T) geometry[T] {
	return geometry[T]{value: value}
}

func (g geometry[T]) AppendAny(p []byte, tag uint8) ([]byte, error) {
	return appendAnyBytesString(p, tag, g.value, contentTypeGeometry), nil
}

func (g geometry[T]) String() string { return string(g.value) }

type json[T []byte | string] struct{ Value T }

func JSON[T []byte | string](value T) json[T] {
	return json[T]{Value: value}
}

func (j json[T]) AppendAny(p []byte, tag uint8) ([]byte, error) {
	return appendAnyBytesString(p, tag, j.Value, contentTypeJSON), nil
}

func (j json[T]) String() string { return string(j.Value) }

func XML[T []byte | string](value T) xml[T] {
	return xml[T]{Value: value}
}

type xml[T []byte | string] struct{ Value T }

func (x xml[T]) AppendAny(p []byte, tag uint8) ([]byte, error) {
	return appendAnyBytesString(p, tag, x.Value, contentTypeXML), nil
}

func (x xml[T]) String() string { return string(x.Value) }
