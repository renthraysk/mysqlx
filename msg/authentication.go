package msg

import (
	"encoding/binary"
	"io"

	"github.com/renthraysk/mysqlx/proto"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
)

const (
	tagAuthenticateStartMechName        = 1
	tagAuthenticateStartAuthData        = 2
	tagAuthenticateStartInitialResponse = 3
)

const (
	tagAuthenticateContinueAuthData = 1
)

// AuthenticateStart builds and sends a AuthenticateStart protobuf message
type AuthenticateStart []byte

// NewAuthenticateStart marshals a AuthenticateStart protobuf message
func NewAuthenticateStart(buf []byte, mechName string) AuthenticateStart {
	b := append(buf[len(buf):], 0, 0, 0, 0, byte(mysqlx.ClientMessages_SESS_AUTHENTICATE_START))
	b = proto.AppendWireString(b, tagAuthenticateStartMechName, mechName)
	return AuthenticateStart(b)
}

// WriteTo writes the AuthenticateStart message to the w, implementation of Msg interface
func (a AuthenticateStart) WriteTo(w io.Writer) (int64, error) {
	binary.LittleEndian.PutUint32(a, uint32(len(a)-4))
	n, err := w.Write(a)
	return int64(n), err
}

// SetAuthData sets the optional authentication data, only used for plain authentication mechanism.
func (a *AuthenticateStart) SetAuthData(authData []byte) {
	*a = proto.AppendWireBytes(*a, tagAuthenticateStartAuthData, authData)
}

// NewAuthenticateContinue marshals AuthenticateContinue protobuf message
func NewAuthenticateContinue(buf []byte, authData []byte) MsgBytes {
	b := append(buf[len(buf):], 0, 0, 0, 0, byte(mysqlx.ClientMessages_SESS_AUTHENTICATE_CONTINUE))
	b = proto.AppendWireBytes(b, tagAuthenticateContinueAuthData, authData)
	return MsgBytes(b)
}
