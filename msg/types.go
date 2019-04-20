package msg

import (
	"github.com/renthraysk/mysqlx/collation"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_resultset"
)

type ContentType uint32

const (
	ContentTypePlain    ContentType = 0
	ContentTypeGeometry             = ContentType(mysqlx_resultset.ContentType_BYTES_GEOMETRY)
	ContentTypeJSON                 = ContentType(mysqlx_resultset.ContentType_BYTES_JSON)
	ContentTypeXML                  = ContentType(mysqlx_resultset.ContentType_BYTES_XML)
)

type String struct {
	Value     string
	Collation collation.Collation
}

func (s *String) AppendArg(a Args) error {
	a.AppendArgString(s.Value, s.Collation)
	return nil
}

type Null struct{}

func (n Null) AppendArg(a Args) error {
	a.AppendArgNull()
	return nil
}

type Geometry []byte

func (g Geometry) AppendArg(a Args) error {
	a.AppendArgBytes(g, ContentTypeGeometry)
	return nil
}

type JSON []byte

func (j JSON) AppendArg(a Args) error {
	a.AppendArgBytes(j, ContentTypeJSON)
	return nil
}

type XML []byte

func (x XML) AppendArg(a Args) error {
	a.AppendArgBytes(x, ContentTypeXML)
	return nil
}
