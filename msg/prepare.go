package msg

import (
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/renthraysk/mysqlx/collation"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_prepare"
	"github.com/renthraysk/mysqlx/slice"
)

const (
	tagPrepareStmtId = 1
	tagPrepareStmt   = 2
)

const (
	tagPrepareOneOfType    = 1
	tagPrepareOneOfExecute = 6
)

func NewPrepare(buf []byte, id uint32, stmt string) Msg {

	n := len(stmt)
	n1 := 1 + SizeUvarint(uint(n)) + n
	n2 := 3 + SizeUvarint(uint(n1)) + n1

	_, b := slice.Allocate(buf, 4+1+1+SizeUvarint32(id)+1+SizeUvarint(uint(n2))+n2)

	b[4] = byte(mysqlx.ClientMessages_PREPARE_PREPARE)

	// Prepare
	b[5] = tagPrepareStmtId<<3 | proto.WireVarint
	i := 6 + binary.PutUvarint(b[6:], uint64(id))
	b[i] = tagPrepareStmt<<3 | proto.WireBytes
	i++
	i += binary.PutUvarint(b[i:], uint64(n2))
	// PrepareOneOf
	b[i] = tagPrepareOneOfType<<3 | proto.WireVarint
	i++
	b[i] = byte(mysqlx_prepare.Prepare_OneOfMessage_STMT)
	i++
	b[i] = tagPrepareOneOfExecute<<3 | proto.WireBytes
	i++
	i += binary.PutUvarint(b[i:], uint64(n1))
	// Execute
	b[i] = tagStmtExecuteStmt<<3 | proto.WireBytes
	i++
	i += binary.PutUvarint(b[i:], uint64(n))
	copy(b[i:], stmt)
	return MsgBytes(b)
}

const (
	tagExecuteStmtId = 1
	tagExecuteArgs   = 2
)

type Execute []byte

func NewExecute(buf []byte, id uint32) Execute {
	n := SizeUvarint32(id)
	_, b := slice.Allocate(buf, 4+1+1+n)

	binary.PutUvarint(b[6:], uint64(id))
	b[4] = byte(mysqlx.ClientMessages_PREPARE_EXECUTE)
	b[5] = tagExecuteStmtId<<3 | proto.WireVarint
	return Execute(b)
}

func (e Execute) WriteTo(w io.Writer) (int64, error) {
	binary.LittleEndian.PutUint32(e, uint32(len(e)-4))
	n, err := w.Write(e)
	return int64(n), err
}

// AppendArgUint appends an uint64 parameter
func (s *Execute) AppendArgUint(v uint64) {
	*s = appendAnyUint(*s, tagExecuteArgs, v)
}

// AppendArgInt appends an int64 parameter
func (s *Execute) AppendArgInt(v int64) {
	*s = appendAnyInt(*s, tagExecuteArgs, v)
}

// AppendArgBytes appends an binary parameter
func (s *Execute) AppendArgBytes(bytes []byte, contentType ContentType) {
	*s = appendAnyBytes(*s, tagExecuteArgs, bytes, contentType)
}

func (s *Execute) AppendArgGeometry(geom []byte) error {
	*s = appendAnyBytes(*s, tagExecuteArgs, geom, ContentTypeGeometry)
	return nil
}

func (s *Execute) AppendArgJSON(json []byte) error {
	*s = appendAnyBytes(*s, tagExecuteArgs, json, ContentTypeJSON)
	return nil
}

func (s *Execute) AppendArgXML(xml []byte) error {
	*s = appendAnyBytes(*s, tagExecuteArgs, xml, ContentTypeXML)
	return nil
}

// AppendArgTime appends a time parameter
func (s *Execute) AppendArgTime(t time.Time) {
	const fmt = "2006-01-02 15:04:05.999999999"
	var b [len(fmt) + 16]byte

	if t.IsZero() {
		s.AppendArgBytes(zeroTime, ContentTypePlain)
		return
	}

	s.AppendArgBytes(t.AppendFormat(b[:0], fmt), ContentTypePlain)
}

// AppendArgString appends a string parameter
func (s *Execute) AppendArgString(str string, collation collation.Collation) {
	*s = appendAnyString(*s, tagExecuteArgs, str, collation)
}

// AppendArgFloat64 appends a float64 parameter
func (s *Execute) AppendArgFloat64(f float64) {
	*s = appendAnyFloat64(*s, tagExecuteArgs, f)
}

// AppendArgFloat32 appends a float32 parameter
func (s *Execute) AppendArgFloat32(f float32) {
	*s = appendAnyFloat32(*s, tagExecuteArgs, f)
}

// AppendArgBool appends a boolean parameter
func (s *Execute) AppendArgBool(b bool) {
	*s = appendAnyBool(*s, tagExecuteArgs, b)
}

func (s *Execute) AppendArgDuration(d time.Duration) error {
	var buf [1 + 20 + 1 + 2 + 1 + 2]byte

	b := buf[:0]
	if d < 0 {
		d = -d
		b = buf[:1]
		b[0] = '-'
	}
	if d > 838*time.Hour+59*time.Minute+59*time.Second {
		return errors.New("time.Duration outside TIME range [-838:59:59, 838:59:59]")
	}

	b = strconv.AppendUint(b, uint64(d/time.Hour), 10)
	i := 2 * (uint(d/time.Minute) % 60)
	j := 2 * (uint(d/time.Second) % 60)
	b = append(b, ':', smallsString[i], smallsString[i+1],
		':', smallsString[j], smallsString[j+1])
	s.AppendArgBytes(b, ContentTypePlain)
	return nil
}

// AppendArgNull appends a NULL parameter
func (s *Execute) AppendArgNull() {
	*s = appendAnyNull(*s, tagExecuteArgs)
}

func (s *Execute) appendArgValue(value interface{}) error {
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
	case time.Duration:
		s.AppendArgDuration(v)
	default:

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

func NewExecuteArgs(buf []byte, id uint32, args []driver.Value) (Msg, error) {
	e := NewExecute(buf, id)
	for i, arg := range args {
		if err := e.appendArgValue(arg); err != nil {
			return nil, errors.Wrapf(err, "unable to serialize argument %d", i)
		}
	}
	return e, nil
}

func NewExecuteNamedArgs(buf []byte, id uint32, args []driver.NamedValue) (Msg, error) {
	e := NewExecute(buf, id)
	for _, arg := range args {
		if len(arg.Name) > 0 {
			return nil, errors.New("mysql does not support the use of named parameters")
		}
		if err := e.appendArgValue(arg.Value); err != nil {
			return nil, errors.Wrapf(err, "unable to serialize named argument %d", arg.Ordinal)
		}
	}
	return e, nil
}

const (
	tagDeallocateStmtID = 1
)

func NewDeallocate(buf []byte, id uint32) MsgBytes {
	_, b := slice.Allocate(buf, 4+1+1+SizeUvarint32(id))
	binary.PutUvarint(b[6:], uint64(id))
	b[4] = byte(mysqlx.ClientMessages_PREPARE_DEALLOCATE)
	b[5] = tagDeallocateStmtID<<3 | proto.WireVarint
	return MsgBytes(b)
}
