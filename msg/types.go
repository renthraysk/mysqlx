package msg

import (
	"github.com/renthraysk/mysqlx/collation"
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
