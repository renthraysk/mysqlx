package mysqlx

import (
	"crypto/tls"
)

func TLSInsecureSkipVerify() *tls.Config {
	return &tls.Config{InsecureSkipVerify: true}
}
