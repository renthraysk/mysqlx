package msg

import (
	"database/sql/driver"
	"encoding/binary"

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
	i := 6 + proto.PutUvarint(b[6:], uint64(id))
	b[i] = tagPrepareStmt<<3 | proto.WireBytes
	i++
	i += proto.PutUvarint(b[i:], uint64(n2))
	// PrepareOneOf
	b[i] = tagPrepareOneOfType<<3 | proto.WireVarint
	i++
	b[i] = byte(mysqlx_prepare.Prepare_OneOfMessage_STMT)
	i++
	b[i] = tagPrepareOneOfExecute<<3 | proto.WireBytes
	i++
	i += proto.PutUvarint(b[i:], uint64(n1))
	// Execute
	b[i] = tagStmtExecuteStmt<<3 | proto.WireBytes
	i++
	i += proto.PutUvarint(b[i:], uint64(n))
	copy(b[i:], stmt)
	return MsgBytes(b)
}

const (
	tagExecuteStmtId = 1
	tagExecuteArgs   = 2
)

func NewExecute(buf []byte, id uint32, args []driver.Value) (MsgBytes, error) {
	b := append(buf[len(buf):], 0, 0, 0, 0, byte(mysqlx.ClientMessages_PREPARE_EXECUTE),
		tagExecuteStmtId<<3|proto.WireVarint, byte(id)|0x80, byte(id>>7)|0x80, byte(id>>14)|0x80, byte(id>>21)|0x80, byte(id>>28))
	i := len(b) - binary.MaxVarintLen32 + proto.SizeVarint32(id)
	b[i-1] &= 0x7F
	b, err := appendAnyValues(b[:i], tagExecuteArgs, args)
	return MsgBytes(b), err
}

func NewExecuteNamed(buf []byte, id uint32, args []driver.NamedValue) (MsgBytes, error) {
	b := append(buf[len(buf):], 0, 0, 0, 0, byte(mysqlx.ClientMessages_PREPARE_EXECUTE),
		tagExecuteStmtId<<3|proto.WireVarint, byte(id)|0x80, byte(id>>7)|0x80, byte(id>>14)|0x80, byte(id>>21)|0x80, byte(id>>28))
	i := len(b) - binary.MaxVarintLen32 + proto.SizeVarint32(id)
	b[i-1] &= 0x7F
	b, err := appendAnyNamedValues(b[:i], tagExecuteArgs, args)
	return MsgBytes(b), err
}

func NewDeallocate(buf []byte, id uint32) MsgBytes {
	const (
		tagDeallocateStmtID = 1
	)
	b := append(buf[len(buf):], 0, 0, 0, 0, byte(mysqlx.ClientMessages_PREPARE_DEALLOCATE),
		tagDeallocateStmtID<<3|proto.WireVarint,
		byte(id)|0x80, byte(id>>7)|0x80, byte(id>>14)|0x80, byte(id>>21)|0x80, byte(id>>28))
	i := len(b) + proto.SizeVarint32(id) - proto.MaxVarintLen32
	b[i-1] &= 0x7F
	return MsgBytes(b[:i])
}
