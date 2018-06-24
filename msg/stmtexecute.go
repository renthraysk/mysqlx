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
	"fmt"
	"io"
	"reflect"
	"time"

	"github.com/renthraysk/mysqlx/collation"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"github.com/renthraysk/mysqlx/slice"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

const (
	tagStmtExecuteStmt            = 1
	tagStmtExecuteArgs            = 2
	tagStmtExecuteNamespace       = 3
	tagStmtExecuteCompactMetadata = 4
)

type ArgAppender interface {
	AppendArg(s *StmtExecute)
}

// StmtExecute is a builder and sender of StmtExecute proto message
type StmtExecute []byte

// NewStmtExecute creates a new StmtExecute which attempts to use the unused capacity of buf.
func NewStmtExecute(buf []byte, stmt string) StmtExecute {
	n := len(stmt)
	i := SizeVarint(uint64(n))
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
	i := SizeVarint(uint64(n))
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

func (s *StmtExecute) AppendArgGeometry(geom []byte) {
	*s = appendAnyBytes(*s, tagStmtExecuteArgs, geom, ContentTypeGeometry)
}

func (s *StmtExecute) AppendArgJSON(json []byte) {
	*s = appendAnyBytes(*s, tagStmtExecuteArgs, json, ContentTypeJSON)
}

func (s *StmtExecute) AppendArgXML(xml []byte) {
	*s = appendAnyBytes(*s, tagStmtExecuteArgs, xml, ContentTypeXML)
}

var zeroTime = []byte{'0', '0', '0', '0', '-', '0', '0', '-', '0', '0'}

// AppendArgTime appends a time parameter
func (s *StmtExecute) AppendArgTime(t time.Time) {
	const fmt = "2006-01-02 15:04:05.999999999"
	var b [len(fmt) + 16]byte

	if t.IsZero() {
		s.AppendArgBytes(zeroTime, ContentTypePlain)
		return
	}

	s.AppendArgBytes(t.AppendFormat(b[:0], fmt), ContentTypePlain)
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

func (s *StmtExecute) appendArgValue(value interface{}) error {
	if value == nil {
		s.AppendArgNull()
		return nil
	}
	switch v := value.(type) {
	case string:
		s.AppendArgString(v, 0)
	case []byte:
		s.AppendArgBytes(v, ContentTypePlain)
	case uint:
		s.AppendArgUint(uint64(v))
	case uint8:
		s.AppendArgUint(uint64(v))
	case uint16:
		s.AppendArgUint(uint64(v))
	case uint32:
		s.AppendArgUint(uint64(v))
	case uint64:
		s.AppendArgUint(v)
	case int:
		s.AppendArgInt(int64(v))
	case int8:
		s.AppendArgInt(int64(v))
	case int16:
		s.AppendArgInt(int64(v))
	case int32:
		s.AppendArgInt(int64(v))
	case int64:
		s.AppendArgInt(v)
	case bool:
		s.AppendArgBool(v)
	case float32:
		s.AppendArgFloat32(v)
	case float64:
		s.AppendArgFloat64(v)
	case time.Time:
		s.AppendArgTime(v)

	default:
		if a, ok := v.(ArgAppender); ok {
			a.AppendArg(s)
			return nil
		}

		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Ptr:
			if rv.IsNil() {
				s.AppendArgNull()
				return nil
			}
			return s.appendArgValue(rv.Elem().Interface())
		default:
			return fmt.Errorf("unsupported type %T, a %s", value, rv.Kind())
		}
	}
	return nil
}

// StmtValues serialises a SQL statement and arguments into a Msg for sending to MySQL.
func StmtValues(buf []byte, stmt string, args []driver.Value) (Msg, error) {
	s := NewStmtExecute(buf, stmt)
	for i, arg := range args {
		if err := s.appendArgValue(arg); err != nil {
			return nil, errors.Wrapf(err, "unable to serialize argument %d", i)
		}
	}
	return s, nil
}

// StmtNamedValues serialises a SQL statement and named arguments into a Msg for sending to MySQL.
// Named arguments are not supported by MySQL, and will result in a error.
func StmtNamedValues(buf []byte, stmt string, args []driver.NamedValue) (Msg, error) {
	s := NewStmtExecute(buf, stmt)
	for _, arg := range args {
		if len(arg.Name) > 0 {
			return nil, errors.New("mysql does not support the use of Named Parameters")
		}
		if err := s.appendArgValue(arg.Value); err != nil {
			return nil, errors.Wrapf(err, "unable to serialize named argument %d", arg.Ordinal)
		}
	}
	return s, nil
}
