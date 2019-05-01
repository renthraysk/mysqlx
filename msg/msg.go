package msg

import (
	"encoding/binary"
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"github.com/renthraysk/mysqlx/slice"
)

const (
	headerSize = 5
)

// Msg generic interface for client to server messages.
type Msg interface {
	io.WriterTo
}

// MsgBytes generic implementation of Msg, for simple byte slices
type MsgBytes []byte

func (m MsgBytes) WriteTo(w io.Writer) (int64, error) {
	binary.LittleEndian.PutUint32(m, uint32(len(m)-4))
	n, err := w.Write(m)
	return int64(n), err
}

// ConnectionClose appends the client close message to buf, and returns Msg to send to server
func ConnectionClose(buf []byte) MsgBytes {
	_, b := slice.Allocate(buf, headerSize)
	b[4] = byte(mysqlx.ClientMessages_CON_CLOSE)
	return MsgBytes(b)
}

// SessionReset appends the client session reset message to buf, and returns Msg to send to server
func SessionReset(buf []byte, keepOpen bool) MsgBytes {

	const tagSessionResetKeepOpen = 1

	if keepOpen {
		_, b := slice.Allocate(buf, headerSize+2)
		b[4] = byte(mysqlx.ClientMessages_SESS_RESET)
		b[5] = tagSessionResetKeepOpen<<3 | proto.WireVarint
		b[6] = 1
		return MsgBytes(b)
	}

	_, b := slice.Allocate(buf, headerSize)
	b[4] = byte(mysqlx.ClientMessages_SESS_RESET)
	return MsgBytes(b)
}
