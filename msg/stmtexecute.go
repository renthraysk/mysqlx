package msg

/*
	Byte banging mysql's X Protocol StmtExecute protobuf.

	func FindByName(buf []byte, name string) Msg {
		s := NewStmtExecute(buf, "SELECT id, name, email FROM users WHERE name = ?")
		s.AppendArgString(name)
		return s
	}

	s := FindByName(buf, "Bob")
	_, err := s.WriteTo(conn)			// Run

*/

import (
	"database/sql/driver"
	"encoding/binary"
	"io"

	"github.com/renthraysk/mysqlx/proto"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
)

const xnamespace = "mysqlx"

const (
	tagStmtExecuteStmt            = 1
	tagStmtExecuteArgs            = 2
	tagStmtExecuteNamespace       = 3
	tagStmtExecuteCompactMetadata = 4
)

// StmtExecute is a builder and sender of StmtExecute proto message
type StmtExecute []byte

// NewStmtExecute creates a new StmtExecute which attempts to use the unused capacity of buf.
func NewStmtExecute(buf []byte, stmt string, args []driver.Value) (StmtExecute, error) {
	var err error

	b := append(buf[len(buf):], 0, 0, 0, 0, byte(mysqlx.ClientMessages_SQL_STMT_EXECUTE))
	b = proto.AppendWireString(b, tagStmtExecuteStmt, stmt)
	b, err = appendAnyValues(b, tagStmtExecuteArgs, args)
	return StmtExecute(b), err
}

// NewStmtExecute creates a new StmtExecute which attempts to use the unused capacity of buf.
func NewStmtExecuteNamed(buf []byte, stmt string, args []driver.NamedValue) (StmtExecute, error) {
	var err error

	b := append(buf[len(buf):], 0, 0, 0, 0, byte(mysqlx.ClientMessages_SQL_STMT_EXECUTE))
	b = proto.AppendWireString(b, tagStmtExecuteStmt, stmt)
	b, err = appendAnyNamedValues(b, tagStmtExecuteArgs, args)
	return StmtExecute(b), err
}

// WriteTo writes protobuf marshalled data to w, implementation of Msg interface
func (s StmtExecute) WriteTo(w io.Writer) (int64, error) {
	binary.LittleEndian.PutUint32(s, uint32(len(s)-4))
	n, err := w.Write(s)
	return int64(n), err
}

// SetNamespace serialises the Namespace field of the StmtExecute protobuf.
func (s *StmtExecute) SetNamespace(namespace string) {
	*s = proto.AppendWireString(*s, tagStmtExecuteNamespace, namespace)
}

func Ping(buf []byte) Msg {
	p, _ := NewStmtExecute(buf, "ping", nil)
	p.SetNamespace(xnamespace)
	return p
}
