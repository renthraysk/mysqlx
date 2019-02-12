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

	"github.com/renthraysk/mysqlx/collation"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"github.com/renthraysk/mysqlx/slice"

	"github.com/golang/protobuf/proto"
)

const (
	tagStmtExecuteStmt            = 1
	tagStmtExecuteArgs            = 2
	tagStmtExecuteNamespace       = 3
	tagStmtExecuteCompactMetadata = 4
)

// StmtExecute is a builder and sender of StmtExecute proto message
type StmtExecute []byte

// NewStmtExecute creates a new StmtExecute which attempts to use the unused capacity of buf.
func NewStmtExecute(buf []byte, stmt string) StmtExecute {
	n := len(stmt)
	i := SizeUvarint64(uint64(n))
	buf, b := slice.Allocate(buf, 4+1+1+i+n)
	binary.PutUvarint(b[6:], uint64(n))
	b[4] = byte(mysqlx.ClientMessages_SQL_STMT_EXECUTE)
	b[5] = tagStmtExecuteStmt<<3 | proto.WireBytes
	copy(b[6+i:], stmt)
	return StmtExecute(b)
}

// WriteTo writes protobuf marshalled data to w, implementation of Msg interface
func (s StmtExecute) WriteTo(w io.Writer) (int64, error) {
	binary.LittleEndian.PutUint32(s, uint32(len(s)-4))
	n, err := w.Write(s)
	return int64(n), err
}

// SetNamespace serialises the Namespace field of the StmtExecute protobuf.
func (s *StmtExecute) SetNamespace(namespace string) {
	n := len(namespace)
	i := SizeUvarint64(uint64(n))
	b := *s
	*s, b = slice.ForAppend(b, 1+i+n)
	binary.PutUvarint(b[1:], uint64(n))
	b[0] = tagStmtExecuteNamespace<<3 | proto.WireBytes
	copy(b[1+i:], namespace)
}

// AppendArgUint appends an uint64 parameter
func (s *StmtExecute) AppendArgUint(v uint64) {
	*s = appendAnyUint(*s, tagStmtExecuteArgs, v)
}

// AppendArgInt appends an int64 parameter
func (s *StmtExecute) AppendArgInt(v int64) {
	*s = appendAnyInt(*s, tagStmtExecuteArgs, v)
}

// AppendArgBytes appends an binary parameter
func (s *StmtExecute) AppendArgBytes(bytes []byte, contentType ContentType) {
	*s = appendAnyBytes(*s, tagStmtExecuteArgs, bytes, contentType)
}

// AppendArgString appends a string parameter
func (s *StmtExecute) AppendArgString(str string, collation collation.Collation) {
	*s = appendAnyString(*s, tagStmtExecuteArgs, str, collation)
}

// AppendArgFloat64 appends a float64 parameter
func (s *StmtExecute) AppendArgFloat64(f float64) {
	*s = appendAnyFloat64(*s, tagStmtExecuteArgs, f)
}

// AppendArgFloat32 appends a float32 parameter
func (s *StmtExecute) AppendArgFloat32(f float32) {
	*s = appendAnyFloat32(*s, tagStmtExecuteArgs, f)
}

// AppendArgBool appends a boolean parameter
func (s *StmtExecute) AppendArgBool(b bool) {
	*s = appendAnyBool(*s, tagStmtExecuteArgs, b)
}

// AppendArgNull appends a NULL parameter
func (s *StmtExecute) AppendArgNull() {
	*s = appendAnyNull(*s, tagStmtExecuteArgs)
}

// StmtValues serialises a SQL statement and arguments into a Msg for sending to MySQL.
func StmtValues(buf []byte, stmt string, args []driver.Value) (Msg, error) {
	s := NewStmtExecute(buf, stmt)
	if err := appendArgValues(&s, args); err != nil {
		return nil, err
	}
	return s, nil
}

// StmtNamedValues serialises a SQL statement and named arguments into a Msg for sending to MySQL.
// Named arguments are not supported by MySQL, and will result in a error.
func StmtNamedValues(buf []byte, stmt string, args []driver.NamedValue) (Msg, error) {
	s := NewStmtExecute(buf, stmt)
	if err := appendArgNamedValues(&s, args); err != nil {
		return nil, err
	}
	return s, nil
}
