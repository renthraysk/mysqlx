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

type Null struct{}

func (n Null) AppendAny(p []byte, tag uint8) ([]byte, error) {
	return appendAnyNull(p, tag), nil
}

type Geometry []byte

func (g Geometry) AppendAny(p []byte, tag uint8) ([]byte, error) {
	return appendAnyBytes(p, tag, g, contentTypeGeometry), nil
}

type GeometryString string

func (g GeometryString) AppendAny(p []byte, tag uint8) ([]byte, error) {
	return appendAnyBytesString(p, tag, string(g), contentTypeGeometry), nil
}

type JSON []byte

func (j JSON) AppendAny(p []byte, tag uint8) ([]byte, error) {
	return appendAnyBytes(p, tag, j, contentTypeJSON), nil
}

type JSONString string

func (j JSONString) AppendAny(p []byte, tag uint8) ([]byte, error) {
	return appendAnyBytesString(p, tag, string(j), contentTypeJSON), nil
}

type XML []byte

func (x XML) AppendAny(p []byte, tag uint8) ([]byte, error) {
	return appendAnyBytes(p, tag, x, contentTypeXML), nil
}

type XMLString string

func (x XMLString) AppendAny(p []byte, tag uint8) ([]byte, error) {
	return appendAnyBytesString(p, tag, string(x), contentTypeXML), nil
}
