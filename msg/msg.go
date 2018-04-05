package msg

import (
	"encoding/binary"
	"io"

	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"github.com/renthraysk/mysqlx/slice"
)

const (
	HeaderSize = 5
)

// Msg generic interface for client to server messages.
type Msg interface {
	io.WriterTo
}

// msg generic implementation of Msg, for simple byte slices
type msg []byte

func (m msg) WriteTo(w io.Writer) (int64, error) {
	binary.LittleEndian.PutUint32(m, uint32(len(m)-4))
	n, err := w.Write(m)
	return int64(n), err
}

// ConnectionClose appends the client close message to buf, and returns Msg to send to server
func ConnectionClose(buf []byte) Msg {
	_, b := slice.Allocate(buf, HeaderSize)
	b[4] = byte(mysqlx.ClientMessages_CON_CLOSE)
	return msg(b)
}

// SessionReset appends the client session reset message to buf, and returns Msg to send to server
func SessionReset(buf []byte) Msg {
	_, b := slice.Allocate(buf, HeaderSize)
	b[4] = byte(mysqlx.ClientMessages_SESS_RESET)
	return msg(b)
}
