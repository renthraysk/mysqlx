package mysqlx

import (
	"encoding/binary"
	"fmt"
	"time"
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

func unmarshalTime(b []byte) (interface{}, error) {

	var min, sec, usec uint64

	i := 1
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
	if j > 0 {
		return nil, fmt.Errorf("failed to decode time (%x)", b)
	}

	d := hour * uint64(time.Hour)
	d += min * uint64(time.Minute)
	d += sec * uint64(time.Second)
	d += usec * uint64(time.Microsecond)

	if b[0] == 0x01 {
		return time.Duration(-int64(d)), nil
	}

	return time.Duration(d), nil
}
