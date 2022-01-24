package sha256

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/renthraysk/mysqlx/authentication"
	"github.com/renthraysk/mysqlx/msg"
	"github.com/renthraysk/mysqlx/slice"
)

type Auth struct{}

// New returns an implementation of authentication.StartContinuer using the sha256_memory authentication mechanism
func New() *Auth {
	return &Auth{}
}

func (Auth) Start(buf []byte, credentials authentication.Credentials) msg.MsgBytes {
	return msg.NewAuthenticateStart(buf, "SHA256_MEMORY", nil)
}

func (Auth) Continue(buf []byte, credentials authentication.Credentials, authData []byte) msg.MsgBytes {

	n := len(credentials.Database()) + 1 + len(credentials.UserName()) + 1 + hex.EncodedLen(sha256.Size)

	// Slice off some bytes for computing the authentication data
	buf, ad := slice.Allocate(buf, n)

	i := copy(ad, credentials.Database())
	ad[i] = 0
	i++
	i += copy(ad[i:], credentials.UserName())
	ad[i] = 0
	i++

	h := sha256.New()
	h.Write([]byte(credentials.Password()))
	h1 := h.Sum(ad[i:i])

	h.Reset()
	h.Write(h1)
	h2 := h.Sum(ad[i+sha256.Size : i+sha256.Size])

	h.Reset()
	h.Write(h2)
	h.Write(authData)
	h2 = h.Sum(h2[:0])

	for i, x := range h1 {
		h2[i] ^= x
	}
	hex.Encode(ad[i:], h2)

	return msg.NewAuthenticateContinue(buf, ad)
}
