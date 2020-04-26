package msg

import (
	"database/sql/driver"
	"encoding/binary"
	"io"

	"github.com/renthraysk/mysqlx/collation"
	"github.com/renthraysk/mysqlx/proto"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_prepare"
	"github.com/renthraysk/mysqlx/slice"
)

func NewPrepare(buf []byte, id uint32, stmt string) Msg {

	const (
		tagPrepareStmtId = 1
		tagPrepareStmt   = 2
	)

	const (
		tagPrepareOneOfType    = 1
		tagPrepareOneOfExecute = 6
	)

	n := len(stmt)
	n1 := 1 + proto.SizeVarint(uint(n)) + n
	n2 := 3 + proto.SizeVarint(uint(n1)) + n1

	_, b := slice.Allocate(buf, 4+1+1+proto.SizeVarint32(id)+1+proto.SizeVarint(uint(n2))+n2)

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
	n := proto.SizeVarint32(id)
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

// AppendArgBytes appends an binary parameter
func (s *Execute) AppendArgBytesString(str string, contentType ContentType) {
	*s = appendAnyBytesString(*s, tagExecuteArgs, str, contentType)
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

// AppendArgNull appends a NULL parameter
func (s *Execute) AppendArgNull() {
	*s = appendAnyNull(*s, tagExecuteArgs)
}

func NewExecuteArgs(buf []byte, id uint32, args []driver.Value) (Msg, error) {
	e := NewExecute(buf, id)
	if err := appendArgValues(&e, args); err != nil {
		return nil, err
	}
	return e, nil
}

func NewExecuteNamedArgs(buf []byte, id uint32, args []driver.NamedValue) (Msg, error) {
	e := NewExecute(buf, id)
	if err := appendArgNamedValues(&e, args); err != nil {
		return nil, err
	}
	return e, nil
}

const (
	tagDeallocateStmtID = 1
)

func NewDeallocate(buf []byte, id uint32) MsgBytes {
	_, b := slice.Allocate(buf, 4+1+1+proto.SizeVarint32(id))
	binary.PutUvarint(b[6:], uint64(id))
	b[4] = byte(mysqlx.ClientMessages_PREPARE_DEALLOCATE)
	b[5] = tagDeallocateStmtID<<3 | proto.WireVarint
	return MsgBytes(b)
}
