package plain

import (
	"github.com/renthraysk/mysqlx/authentication"
	"github.com/renthraysk/mysqlx/msg"
	"github.com/renthraysk/mysqlx/slice"
)

type Auth struct{}

// New returns an implementation of authentication.StartContinuer using the mysql plain authentication mechanism
func New() *Auth {
	return &Auth{}
}

func (Auth) Start(buf []byte, credentials authentication.Credentials) msg.AuthenticateStart {
	n := len(credentials.Database()) + 1 + len(credentials.UserName()) + 1 + len(credentials.Password())

	buf, ad := slice.Allocate(buf, n)

	i := copy(ad, credentials.Database())
	ad[i] = 0
	i++
	i += copy(ad[i:], credentials.UserName())
	ad[i] = 0
	i++
	copy(ad[i:], credentials.Password())

	m := msg.NewAuthenticateStart(buf, "PLAIN")
	m.SetAuthData(ad)
	return m
}
