package mysql41

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/renthraysk/mysqlx/authentication"
	"github.com/renthraysk/mysqlx/msg"
	"github.com/renthraysk/mysqlx/slice"
)

type Auth struct{}

// New returns an implementation of authentication.StartContinuer using the mysql41 password authentication mechanism
func New() *Auth {
	return &Auth{}
}

func (Auth) Start(buf []byte, credentials authentication.Credentials) msg.MsgBytes {
	return msg.NewAuthenticateStart(buf, "MYSQL41", nil)
}

func (Auth) Continue(buf []byte, credentials authentication.Credentials, authData []byte) msg.MsgBytes {
	n := len(credentials.Database()) + 1 + len(credentials.UserName()) + 1
	if len(credentials.Password()) > 0 {
		n += 1 + 2*sha1.Size
	}
	// Slice off some bytes for computing the authentication data
	buf, ad := slice.Allocate(buf, n)

	i := copy(ad, credentials.Database())
	ad[i] = 0
	i++
	i += copy(ad[i:], credentials.UserName())
	ad[i] = 0
	i++
	if len(credentials.Password()) > 0 {
		ad[i] = '*'
		i++

		h1 := ad[i : i+sha1.Size]
		h2 := ad[i+sha1.Size:]

		h := sha1.New()
		h.Write([]byte(credentials.Password()))
		h.Sum(h1[:0])

		h.Reset()
		h.Write(h1)
		h.Sum(h2[:0])

		h.Reset()
		h.Write(authData)
		h.Write(h2)
		h.Sum(h2[:0])

		for i, x := range h1 {
			h2[i] ^= x
		}
		hex.Encode(ad[i:], h2)
	}

	return msg.NewAuthenticateContinue(buf, ad)
}
