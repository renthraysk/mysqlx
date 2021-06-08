package msg

import (
	"encoding/binary"
	"io"

	"github.com/renthraysk/mysqlx/proto"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
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
	return MsgBytes(append(buf[len(buf):], 0, 0, 0, 0, byte(mysqlx.ClientMessages_CON_CLOSE)))
}

// SessionReset appends the client session reset message to buf, and returns Msg to send to server
func SessionReset(buf []byte, keepOpen bool) MsgBytes {
	const tagSessionResetKeepOpen = 1

	if keepOpen {
		b := append(buf[len(buf):], 0, 0, 0, 0, byte(mysqlx.ClientMessages_SESS_RESET),
			tagSessionResetKeepOpen<<3|proto.WireVarint, 1)
		return MsgBytes(b)
	}
	return MsgBytes(append(buf[len(buf):], 0, 0, 0, 0, byte(mysqlx.ClientMessages_SESS_RESET)))
}
