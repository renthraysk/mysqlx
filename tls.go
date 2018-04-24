package mysqlx

import (
	"crypto/tls"
)

// TLSInsecureSkipVerify returns a new tls.Config{} with InsecureSkipVerify set.
func TLSInsecureSkipVerify() *tls.Config {
	return &tls.Config{InsecureSkipVerify: true}
}
