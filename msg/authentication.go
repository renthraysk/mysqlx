package msg

import (
	"github.com/renthraysk/mysqlx/proto"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
)

// NewAuthenticateStart marshals a AuthenticateStart protobuf message
func NewAuthenticateStart(buf []byte, mechName string, authData []byte) MsgBytes {
	const (
		tagAuthenticateStartMechName        = 1
		tagAuthenticateStartAuthData        = 2
		tagAuthenticateStartInitialResponse = 3
	)

	b := append(buf[len(buf):], 0, 0, 0, 0, byte(mysqlx.ClientMessages_SESS_AUTHENTICATE_START))
	b = proto.AppendWireString(b, tagAuthenticateStartMechName, mechName)
	if len(authData) > 0 {
		b = proto.AppendWireBytes(b, tagAuthenticateStartAuthData, authData)
	}
	return MsgBytes(b)
}

// NewAuthenticateContinue marshals AuthenticateContinue protobuf message
func NewAuthenticateContinue(buf []byte, authData []byte) MsgBytes {
	const (
		tagAuthenticateContinueAuthData = 1
	)

	b := append(buf[len(buf):], 0, 0, 0, 0, byte(mysqlx.ClientMessages_SESS_AUTHENTICATE_CONTINUE))
	b = proto.AppendWireBytes(b, tagAuthenticateContinueAuthData, authData)
	return MsgBytes(b)
}
