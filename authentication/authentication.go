package authentication

import "github.com/renthraysk/mysqlx/msg"

type Credentials interface {
	UserName() string
	Password() string
	Database() string
}

type Starter interface {
	Start(buf []byte, credentials Credentials) msg.AuthenticateStart
}

type StartContinuer interface {
	Starter
	Continue(buf []byte, credentials Credentials, authData []byte) msg.MsgBytes
}
