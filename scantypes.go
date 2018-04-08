package mysqlx

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

type Set struct {
	Members []string
	Valid   bool
}

func (s *Set) Scan(src interface{}) error {
	s.Members = s.Members[:0]
	s.Valid = false
	if src == nil {
		return nil
	}
	switch b := src.(type) {
	case []byte:
		n := uint64(len(b))
		if n == 0 {
			return nil
		}
		s.Valid = true
		str := string(b)
		i := uint64(0)

		for i < n {
			j, nn := binary.Uvarint(b[i:])
			if nn <= 0 {
				return errors.Wrap(io.ErrUnexpectedEOF, "failed to unmarshal set")
			}
			i += uint64(nn)
			j += i
			if j > n {
				j = n
			}
			s.Members = append(s.Members, str[i:j])
			i = j
		}
		return nil
	}
	return fmt.Errorf("Unable to unmarshal type %T to Set", src)
}
