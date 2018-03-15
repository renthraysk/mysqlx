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

type Option func(*Connector) error

type Connector struct {
	network         string
	addr            string
	netDialer       net.Dialer
	tlsConfig       *tls.Config
	database        string
	username        string
	password        string
	authentication  authentication.Starter
	stmtPreparer    StmtPreparer
	sessionResetter SessionResetter

	bufferSize int
}

const minBufferSize = 32 * 1024

// authentication.Credentials interface implementation
func (cnn *Connector) UserName() string { return cnn.username }
func (cnn *Connector) Password() string { return cnn.password }
func (cnn *Connector) Database() string { return cnn.database }

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

func WithDialer(netDialer net.Dialer) Option {
	return func(cnn *Connector) error {
		cnn.netDialer = netDialer
		return nil
	}
}

func WithDatabase(database string) Option {
	return func(cnn *Connector) error {
		cnn.database = database
		return nil
	}
}

func WithAuthentication(auth authentication.Starter) Option {
	return func(cnn *Connector) error {
		cnn.authentication = auth
		return nil
	}
}

func WithUserPassword(username, password string) Option {
	return func(cnn *Connector) error {
		cnn.username = username
		cnn.password = password
		return nil
	}
}

func WithTLSConfig(tlsConfig *tls.Config) Option {
	return func(cnn *Connector) error {
		cnn.tlsConfig = tlsConfig
		return nil
	}
}

func WithBufferSize(size int) Option {
	return func(cnn *Connector) error {
		if size > minBufferSize {
			cnn.bufferSize = size
		}
		return nil
	}
}

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
			if _, err := conn.ExecMsg(ctx, s); err != nil {
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

func (cnn *Connector) Driver() driver.Driver {
	return nil
}
