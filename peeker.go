package mysqlx

import (
	"fmt"
	"io"
	"math"
	"time"
)

const (
	ErrNegativeCount = errorString("negative count")
	ErrAdvanceTooFar = errorString("advance too far")
)

const (
	minReadSize   = 512
	readAheadSize = 256
)

type Reader interface {
	io.Reader
	SetReadDeadline(time.Time) error
}

type Peeker struct {
	r    Reader
	pbuf []byte // slice of available bytes to peek from, references into rbuf below
	rbuf []byte // underlying Read() buffer
}

func NewPeeker(r Reader) *Peeker {
	return &Peeker{r: r}
}

func (p *Peeker) Reset(r Reader) {
	p.r = r
	p.pbuf = p.rbuf[:0:0]
}

// buf returns a suitable sized slice to read atleast n bytes into
func (p *Peeker) buf(n int) []byte {
	nn := minReadSize
	if n > nn {
		nn = n
		// overflow checking
		if nn < math.MaxInt32-readAheadSize {
			nn += readAheadSize
		}
	}
	if n > len(p.rbuf) {
		p.rbuf = make([]byte, nn)
	} else if len(p.rbuf) > nn {
		// Limit how much to read as likely have to copy it
		return p.rbuf[:nn]
	}
	return p.rbuf
}

func (p *Peeker) fill(n int, deadline time.Time) ([]byte, error) {
	buf := p.buf(n)
	i := copy(buf, p.pbuf)
	if err := p.r.SetReadDeadline(deadline); err != nil {
		return nil, fmt.Errorf("failed to set read deadline: %w", err)
	}
	r, err := p.r.Read(buf[i:])
	i += r
	for i < n && err == nil {
		r, err = p.r.Read(buf[i:])
		i += r
	}
	buf = buf[:i:i]
	p.pbuf = buf
	if n > i {
		return buf, err
	}
	return buf[:n:n], nil
}

func (p *Peeker) Peek(n int, deadline time.Time) ([]byte, error) {
	if n < 0 {
		return nil, ErrNegativeCount
	}
	if n > len(p.pbuf) {
		return p.fill(n, deadline)
	}
	return p.pbuf[:n:n], nil
}

func (p *Peeker) Discard(n int) error {
	if n < 0 {
		return ErrNegativeCount
	}
	if n > len(p.pbuf) {
		return ErrAdvanceTooFar
	}
	p.pbuf = p.pbuf[n:]
	return nil
}
