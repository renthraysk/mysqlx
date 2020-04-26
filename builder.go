package mysqlx

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"

	"github.com/renthraysk/mysqlx/msg"
	"github.com/renthraysk/mysqlx/proto"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_expect"
)

type builder struct {
	buf *bytes.Buffer
	tmp [64]byte
}

func newBuilder() *builder {
	return &builder{
		buf: &bytes.Buffer{},
	}
}

func newBuilderBuffer(buf []byte) *builder {
	return &builder{
		buf: bytes.NewBuffer(buf),
	}
}

type onError byte

const (
	onErrorContinue onError = '0'
	onErrorFail     onError = '1'
)

const (
	tagOpenOperation = 1
	tagOpenCondition = 2
)
const (
	tagConditionKey       = 1
	tagConditionValue     = 2
	tagConditionOperation = 3
)

func (b *builder) Reset() {
	b.buf.Reset()
}

func (b *builder) WriteExpectOpen(onError onError) error {

	s := append(b.tmp[:4], byte(mysqlx.ClientMessages_EXPECT_OPEN),
		tagOpenOperation<<3|proto.WireVarint, byte(mysqlx_expect.Open_EXPECT_CTX_EMPTY),
		tagOpenCondition<<3|proto.WireBytes, 5,
		tagConditionKey<<3|proto.WireVarint, byte(mysqlx_expect.Open_Condition_EXPECT_NO_ERROR),
		tagConditionValue<<3|proto.WireBytes, 1, byte(onError),
	)
	binary.LittleEndian.PutUint32(s, uint32(len(s)-4))

	_, err := b.buf.Write(s)
	return err
}

func (b *builder) WriteExpectField(field string) error {
	n := len(field)
	i := 8 + proto.PutUvarint(b.tmp[8:], uint64(2+1+proto.SizeVarint(uint(n))+n))
	b.tmp[4] = byte(mysqlx.ClientMessages_EXPECT_OPEN)
	b.tmp[5] = tagOpenOperation<<3 | proto.WireVarint
	b.tmp[6] = byte(mysqlx_expect.Open_EXPECT_CTX_EMPTY)
	b.tmp[7] = tagOpenCondition<<3 | proto.WireBytes

	b.tmp[i] = tagConditionKey<<3 | proto.WireVarint
	i++
	b.tmp[i] = byte(mysqlx_expect.Open_Condition_EXPECT_FIELD_EXIST)
	i++
	b.tmp[i] = tagConditionValue<<3 | proto.WireBytes
	i++
	i += proto.PutUvarint(b.tmp[i:], uint64(n))
	binary.LittleEndian.PutUint32(b.tmp[:], uint32(i+n-4))
	_, err := b.buf.Write(b.tmp[:i])
	if err == nil {
		_, err = b.buf.WriteString(field)
	}
	return err
}

func (b *builder) WriteExpectClose() error {
	binary.LittleEndian.PutUint32(b.tmp[:], 1)
	b.tmp[4] = byte(mysqlx.ClientMessages_EXPECT_CLOSE)
	_, err := b.buf.Write(b.tmp[:5])
	return err
}

func (b *builder) WriteSessionReset(keepOpen bool) error {
	s := msg.SessionReset(b.tmp[:0], keepOpen)
	_, err := s.WriteTo(b.buf)
	return err
}

func (b *builder) WriteStmtExecute(stmt string, args ...driver.Value) error {
	s, err := msg.NewStmtExecute(b.tmp[:0], stmt, args)
	if err == nil {
		_, err = s.WriteTo(b.buf)
	}
	return err
}

func (b *builder) WritePing() error {
	p := msg.Ping(b.tmp[:0])
	_, err := p.WriteTo(b.buf)
	return err
}

func (b *builder) Bytes() []byte {
	return b.buf.Bytes()
}
