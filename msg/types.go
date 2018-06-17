package msg

import (
	"github.com/renthraysk/mysqlx/collation"
)

type String struct {
	Value     string
	Collation collation.Collation
}

func (s *String) AppendArg(se *StmtExecute) {
	se.AppendArgString(s.Value, s.Collation)
}

type Null struct{}

func (n Null) AppendArg(s *StmtExecute) {
	s.AppendArgNull()
}

type Geometry []byte

func (g Geometry) AppendArg(s *StmtExecute) {
	s.AppendArgGeometry(g)
}

type JSON []byte

func (j JSON) AppendArg(s *StmtExecute) {
	s.AppendArgJSON(j)
}

type XML []byte

func (x XML) AppendArg(s *StmtExecute) {
	s.AppendArgXML(x)
}
