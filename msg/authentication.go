package msg

import (
	"encoding/binary"
	"io"

	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"github.com/renthraysk/mysqlx/slice"

	"github.com/golang/protobuf/proto"
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
	n := len(mechName)
	i := SizeVarint(uint64(n))
	buf, b := slice.Allocate(buf, 4+1+1+i+n)
	binary.PutUvarint(b[6:], uint64(n))
	b[4] = byte(mysqlx.ClientMessages_SESS_AUTHENTICATE_START)
	b[5] = tagAuthenticateStartMechName<<3 | proto.WireBytes
	copy(b[6+i:], mechName)
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
	n := len(authData)
	i := SizeVarint(uint64(n))
	b := *a
	*a, b = slice.ForAppend(b, 1+i+n)
	binary.PutUvarint(b[1:], uint64(n))
	b[0] = tagAuthenticateStartAuthData<<3 | proto.WireBytes
	copy(b[1+i:], authData)
}

// NewAuthenticateContinue marshals AuthenticateContinue protobuf message
func NewAuthenticateContinue(buf []byte, authData []byte) Msg {
	n := len(authData)
	i := SizeVarint(uint64(n))
	buf, b := slice.Allocate(buf, 4+1+1+i+n)
	binary.PutUvarint(b[6:], uint64(n))
	b[4] = byte(mysqlx.ClientMessages_SESS_AUTHENTICATE_CONTINUE)
	b[5] = tagAuthenticateContinueAuthData<<3 | proto.WireBytes
	copy(b[6+i:], authData)
	return msg(b)
}
