package mysqlx

import (
	"context"
	"crypto/tls"
	"database/sql/driver"
	"net"

	"github.com/renthraysk/mysqlx/authentication"
	"github.com/renthraysk/mysqlx/authentication/native"
	"github.com/renthraysk/mysqlx/msg"

	"github.com/pkg/errors"
)

// Option is a functional option for creating the Connector
type Option func(*Connector) error

// Connector is the database/sql.Connector implementation
type Connector struct {
	network         string
	addr            string
	netDialer       net.Dialer
	tlsConfig       *tls.Config
	database        string
	username        string
	password        string
	authentication  authentication.Starter
	stmtPreparer    stmtPreparer
	sessionResetter sessionResetter

	bufferSize int
}

const minBufferSize = 32 * 1024

// UserName returns the user name of the account to authenticate with. Part of authentication.Credentials inteface.
func (cnn *Connector) UserName() string { return cnn.username }

// Password returns the password of the account to authenticate with. Part of authentication.Credentials inteface.
func (cnn *Connector) Password() string { return cnn.password }

// Database returns the database name to authenticate with. Part of authentication.Credentials inteface.
func (cnn *Connector) Database() string { return cnn.database }

// New creates a database/sql.Connector
func New(network, addr string, options ...Option) (*Connector, error) {

	cnn := &Connector{
		network:         network,
		addr:            addr,
		authentication:  native.New(),
		stmtPreparer:    noStmtPreparer,
		sessionResetter: noSessionResetter,
		bufferSize:      minBufferSize,
	}

	for _, opt := range options {
		if err := opt(cnn); err != nil {
			return nil, err
		}
	}

	return cnn, nil
}

// WithDialer replaces the default net.Dialer for connecting to mysql.
func WithDialer(netDialer net.Dialer) Option {
	return func(cnn *Connector) error {
		cnn.netDialer = netDialer
		return nil
	}
}

// WithDatabase sets the database the connector will be default after successful connection and authentication
func WithDatabase(database string) Option {
	return func(cnn *Connector) error {
		cnn.database = database
		return nil
	}
}

// WithAuthentication set the authentication mechanism that will authentication with.
// If authenticating a connection over TLS then either authentication/native for accounts using the mysql_native_password authentication plugin or authentication/sha256 for those using sha256_password or caching_sha2_password.
// If not using a TLS connection then authentication/native is the only reliable option.
func WithAuthentication(auth authentication.Starter) Option {
	return func(cnn *Connector) error {
		cnn.authentication = auth
		return nil
	}
}

// WithUserPassword set the username and password pair of the account to authenticate with.
func WithUserPassword(username, password string) Option {
	return func(cnn *Connector) error {
		cnn.username = username
		cnn.password = password
		return nil
	}
}

// WithTLSConfig set the TLS configuration to connect to mysqlx with.
func WithTLSConfig(tlsConfig *tls.Config) Option {
	return func(cnn *Connector) error {
		cnn.tlsConfig = tlsConfig
		return nil
	}
}

// WithBufferSize sets the internal read/write buffer size. It will be automatically enlarged if larger reads are required.
func WithBufferSize(size int) Option {
	return func(cnn *Connector) error {
		if size > minBufferSize {
			cnn.bufferSize = size
		}
		return nil
	}
}

// Connect is the database/sql.Connector Connect() implementation
func (cnn *Connector) Connect(ctx context.Context) (driver.Conn, error) {
	return cnn.connect(ctx)
}

func (cnn *Connector) connect(ctx context.Context) (Conn, error) {
	netConn, err := cnn.netDialer.DialContext(ctx, cnn.network, cnn.addr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial")
	}

	conn := &conn{
		netConn:   netConn,
		connector: cnn,
		buf:       make([]byte, cnn.bufferSize),
	}

	if tc, ok := netConn.(*net.TCPConn); ok {
		if err := tc.SetKeepAlive(true); err != nil {
			netConn.Close()
			return nil, errors.Wrap(err, "failed to set keep alive")
		}

		// TLS
		if cnn.tlsConfig != nil {
			s := msg.NewCapabilitySetTLSEnable(conn.buf[:0])
			if _, err := conn.execMsg(ctx, s); err != nil {
				return nil, errors.Wrap(err, "failed to set TLS capability")
			}
			tlsConn := tls.Client(conn.netConn, cnn.tlsConfig)
			if err := tlsConn.Handshake(); err != nil {
				tlsConn.Close()
				return nil, errors.Wrap(err, "failed TLS handshake")
			}
			conn.netConn = tlsConn
		}
	}

	if err := conn.authenticate(ctx, cnn.authentication); err != nil {
		return nil, errors.Wrap(err, "failed to authenticate")
	}
	return conn, nil
}

// Driver is the database/sql.Connector Driver() implementation
func (cnn *Connector) Driver() driver.Driver {
	return nil
}
