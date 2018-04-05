package mysqlx

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

func unmarshalDateTime(b []byte) (time.Time, error) {
	year, i := binary.Uvarint(b)
	if i <= 0 {
		return time.Time{}, fmt.Errorf("failed to decode datetime year (%x)", b)
	}
	month, j := binary.Uvarint(b[i:])
	if j <= 0 {
		return time.Time{}, fmt.Errorf("failed to decode datetime month (%x)", b)
	}
	i += j
	day, j := binary.Uvarint(b[i:])
	if j <= 0 {
		return time.Time{}, fmt.Errorf("failed to decode datetime day (%x)", b)
	}
	i += j

	var min, sec, usec uint64

	hour, j := binary.Uvarint(b[i:])
	if j > 0 {
		i += j
		min, j = binary.Uvarint(b[i:])
		if j > 0 {
			i += j
			sec, j = binary.Uvarint(b[i:])
			if j > 0 {
				i += j
				usec, j = binary.Uvarint(b[i:])
			}
		}
	}
	if j < 0 {
		return time.Time{}, fmt.Errorf("failed to decode datetime time (%x)", b)
	}
	return time.Date(int(year), time.Month(month), int(day), int(hour), int(min), int(sec), int(usec)*1000, time.UTC), nil
}

func unmarshalDecimal(b []byte) (string, error) {
	if len(b) < 2 {
		return "", fmt.Errorf("failed to parse decimal %#v", b)
	}

	var h uint8

	buf := [96]byte{'-'} // assume negative, easier to slice off if non -'ve
	r := buf[:1]

	for _, l := range b[1:] {
		h = l >> 4
		if h > 9 {
			break
		}
		l &= 0x0F
		if l > 9 {
			r = append(r, '0'+h)
			h = l
			break
		}
		r = append(r, '0'+h, '0'+l)
	}

	// If not negative remove the premptive -
	if h != 0x0B && h != 0x0D {
		r = r[1:]
	}

	if s := b[0]; s > 0 {
		i := len(r) - int(s)
		if i < 0 {
			return "", fmt.Errorf("scale (%d) exceeds precision (%d) in %#v", s, len(r), b)
		}
		r = append(r, 0)
		copy(r[i+1:], r[i:])
		r[i] = '.'
	}
	return string(r), nil
}

func unmarshalSet(b []byte) ([]string, error) {
	return nil, errors.New("unimplemented SET unmarshalling")
}

func unmarshalTime(b []byte) (interface{}, error) {
	return nil, errors.New("unimplemnted TIME unmarshalling")
}

func unmarshalBit(b []byte) (interface{}, error) {
	return nil, errors.New("unimplemnted BIT unmarshalling")
}
