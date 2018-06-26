package msg

import (
	"github.com/renthraysk/mysqlx/collation"
)

type String struct {
	Value     string
	Collation collation.Collation
}

func (s *String) AppendArg(se *StmtExecute) error {
	se.AppendArgString(s.Value, s.Collation)
	return nil
}

type Null struct{}

func (n Null) AppendArg(s *StmtExecute) error {
	s.AppendArgNull()
	return nil
}

type Geometry []byte

func (g Geometry) AppendArg(s *StmtExecute) error {
	return s.AppendArgGeometry(g)
}

type JSON []byte

func (j JSON) AppendArg(s *StmtExecute) error {
	return s.AppendArgJSON(j)
}

type XML []byte

func (x XML) AppendArg(s *StmtExecute) error {
	return s.AppendArgXML(x)
}
