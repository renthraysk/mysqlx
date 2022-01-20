package mysqlx

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/renthraysk/mysqlx/authentication"
	"github.com/renthraysk/mysqlx/authentication/mysql41"
	"github.com/renthraysk/mysqlx/errs"
	"github.com/renthraysk/mysqlx/msg"
)

// Dailer interface documenting our requirements for dialing a MySQL server.
type Dialer interface {
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
}

// Connector is the database/sql.Connector implementation
type Connector struct {
	network        string
	addr           string
	dialer         Dialer
	tlsConfig      *tls.Config
	database       string
	username       string
	password       string
	authentication authentication.Starter
	bufferSize     int

	resetBuild    sync.Once
	resetKeepOpen bool
	resetFuncs    []func(b *builder) error
	reset         []byte

	connectAttrs map[string]string
}

// Option is a functional option for creating the Connector
type Option func(*Connector) error

const minBufferSize = 4 * 1024

// UserName returns the user name of the account to authenticate with. Part of authentication.Credentials interface.
func (cnn *Connector) UserName() string { return cnn.username }

// Password returns the password of the account to authenticate with. Part of authentication.Credentials interface.
func (cnn *Connector) Password() string { return cnn.password }

// Database returns the database name to authenticate with. Part of authentication.Credentials interface.
func (cnn *Connector) Database() string { return cnn.database }

// New creates a database/sql.Connector
func New(network, addr string, options ...Option) (*Connector, error) {
	cnn := &Connector{
		dialer:         new(net.Dialer),
		network:        network,
		addr:           addr,
		authentication: mysql41.New(),
		bufferSize:     minBufferSize,

		resetFuncs: make([]func(b *builder) error, 0, 2),
	}

	for _, opt := range options {
		if err := opt(cnn); err != nil {
			return nil, err
		}
	}
	return cnn, nil
}

// WithDialer replaces the default net.Dialer for connecting to mysql.
func WithDialer(dialer Dialer) Option {
	return func(cnn *Connector) error {
		cnn.dialer = dialer
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
// If authenticating a connection over TLS then either authentication/native or authentication/sha256.
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

// WithBufferSize sets the internal read/write buffer size.
func WithBufferSize(size int) Option {
	return func(cnn *Connector) error {
		if size > minBufferSize {
			cnn.bufferSize = size
		}
		return nil
	}
}

func WithDefaultConnectAttrs() Option {
	attrs := map[string]string{
		"_client_name": "mysqlx",
		"_pid":         strconv.Itoa(os.Getpid()),
		"_platform":    runtime.GOARCH,
		"_os":          runtime.GOOS,
		"program_name": os.Args[0],
	}
	return WithConnectAttrs(attrs)
}

// WithConnectAttrs sets the connection attributes that will be set on connect
func WithConnectAttrs(attrs map[string]string) Option {
	return func(cnn *Connector) error {
		if cnn.connectAttrs == nil {
			cnn.connectAttrs = make(map[string]string, len(attrs))
		}
		for k, v := range attrs {
			cnn.connectAttrs[k] = v
		}
		return nil
	}
}

// WithSQLMode set the default sql_mode
func WithSQLMode(modes ...string) Option {
	return func(cnn *Connector) error {
		cnn.resetFuncs = append(cnn.resetFuncs, func(b *builder) error {
			return b.WriteStmtExecute("SET SESSION sql_mode = ?", strings.Join(modes, ","))
		})
		return nil
	}
}

type SessionVars map[string]any

func (sv SessionVars) build(b *builder) error {

	var s bytes.Buffer

	for k, v := range sv {
		s.Reset()
		s.WriteString("SET SESSION `")
		for i := strings.IndexByte(k, '`'); i >= 0; i = strings.IndexByte(k, '`') {
			i++
			s.WriteString(k[:i])
			s.WriteByte('`')
			k = k[i:]
		}
		s.WriteString(k)
		s.WriteString("` = ?")
		if err := b.WriteStmtExecute(s.String(), v); err != nil {
			return err
		}
	}
	return nil
}

// WithSessionVars sets/resets mysql session variables on connect and every reset.
func WithSessionVars(sv SessionVars) Option {
	return func(cnn *Connector) error {
		cnn.resetFuncs = append(cnn.resetFuncs, sv.build)
		return nil
	}
}

// WithDefaultTxIsolation sets the default transaction isolation level on connect, and at every reset.
func WithDefaultTxIsolation(isolationLevel sql.IsolationLevel) Option {
	set := ""
	switch isolationLevel {
	case sql.LevelDefault:
		return func(cnn *Connector) error { return nil }
	case sql.LevelReadUncommitted:
		set = "SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED"
	case sql.LevelReadCommitted:
		set = "SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED"
	case sql.LevelRepeatableRead:
		set = "SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ"
	case sql.LevelSerializable:
		set = "SET SESSION TRANSACTION ISOLATION LEVEL SERIALIZABLE"
	default:
		return func(cnn *Connector) error {
			return fmt.Errorf("Unsupported default transaction isolation level (%s)", isolationLevel.String())
		}
	}
	return func(cnn *Connector) error {
		cnn.resetFuncs = append(cnn.resetFuncs, func(b *builder) error {
			return b.WriteStmtExecute(set)
		})
		return nil
	}
}

// Connect is the database/sql.Connector Connect() implementation
func (cnn *Connector) Connect(ctx context.Context) (driver.Conn, error) {
	netConn, err := cnn.dialer.DialContext(ctx, cnn.network, cnn.addr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	conn := &conn{
		netConn:   netConn,
		r:         bufio.NewReaderSize(netConn, cnn.bufferSize),
		connector: cnn,
		buf:       make([]byte, cnn.bufferSize),
	}

	// TLS
	if _, ok := netConn.(*net.TCPConn); ok && cnn.tlsConfig != nil {
		s, err := msg.CapabilitySetTLS(conn.buf, true)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal TLS enable CapabilitySet: %w", err)
		}
		if _, err := conn.execMsg(ctx, s); err != nil {
			netConn.Close()
			return nil, fmt.Errorf("failed to set TLS capability: %w", err)
		}
		tlsConn := tls.Client(netConn, cnn.tlsConfig)
		if err := tlsConn.Handshake(); err != nil {
			tlsConn.Close()
			return nil, fmt.Errorf("failed TLS handshake: %w", err)
		}
		conn.netConn = tlsConn
		conn.r.Reset(tlsConn)
	}

	// Connection Attributes
	if len(cnn.connectAttrs) > 0 {
		cs, err := msg.CapabilitySetSessionConnectAttrs(conn.buf, cnn.connectAttrs)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal session connection attributes: %w", err)
		}
		if _, err := conn.execMsg(ctx, cs); err != nil {
			return nil, fmt.Errorf("failed to set session connection attributes: %w", err)
		}
	}

	// Authentication
	if err := conn.authenticate(ctx); err != nil {
		return nil, err
	}

	// Build a byte slice to reset mysql session state every reset.
	// Using err inside this func() to signal err occurring within
	cnn.resetBuild.Do(func() {
		// Build byteslice to be written/executed on every connect & reset.

		const ExpectFieldKeepOpen = "6.1"

		// See if MySQL supports session-reset's keepOpen field...
		b := newBuilder()
		b.WriteExpectField(ExpectFieldKeepOpen)
		b.WriteExpectClose()

		cnn.resetKeepOpen = false
		if err = conn.sendN(ctx, b.Bytes()); err != nil {
			var e *errs.Errors
			if !errors.As(err, &e) {
				return
			}

			if !errs.ErXExpectFieldExistsFailed.Is((*e)[0]) {
				return
			}
			// No session-reset(keep-open) support.
			err = nil
		} else {
			cnn.resetKeepOpen = true
		}

		// Determine if have anything to run per reset
		if !cnn.resetKeepOpen && len(cnn.resetFuncs) == 0 {
			return
		}
		b.Reset()
		b.WriteExpectOpen(onErrorFail)
		// If can perform a session reset without needing to reauthenticate
		// then include session-reset(keepOpen=true) for single write reset.
		if cnn.resetKeepOpen {
			b.WriteSessionReset(true)
		} else {
			// Using a ping as NOP to ensure error indexes remain consistent for tests.
			b.WritePing()
		}
		for _, f := range cnn.resetFuncs {
			if err = f(b); err != nil {
				return
			}
		}
		b.WriteExpectClose()
		cnn.reset = b.Bytes()
	})
	if err != nil {
		return nil, err
	}

	if err := conn.sendN(ctx, cnn.reset); err != nil {
		return nil, err
	}
	return conn, nil
}

// Driver is the database/sql.Connector Driver() implementation
func (cnn *Connector) Driver() driver.Driver {
	return nil
}
