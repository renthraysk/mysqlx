package mysqlx

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type Date struct {
	Year  uint64
	Month uint64
	Day   uint64
}

func (d Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

func (d *Date) Unmarshal(b []byte) error {
	var i, j int

	d.Year, i = binary.Uvarint(b)
	if i <= 0 {
		return errors.Errorf("failed to decode year (%x)", b)
	}
	d.Month, j = binary.Uvarint(b[i:])
	if j <= 0 {
		return errors.Errorf("failed to decode month (%x)", b)
	}
	i += j
	d.Day, j = binary.Uvarint(b[i:])
	if j <= 0 {
		return errors.Errorf("failed to decode day (%x)", b)
	}
	return nil
}

type NullDate struct {
	Date
	Valid bool
}

func (d *NullDate) Scan(src interface{}) error {
	if src == nil {
		d.Valid = false
		return nil
	}
	d.Valid = true
	switch v := src.(type) {
	case []byte:
		return d.Date.Unmarshal(v)
	case Date:
		d.Date = v
		return nil
	}
	return errors.Errorf("unable to convert type %T to %T", src, d)
}

type DateTime struct {
	Year       uint64
	Month      uint64
	Day        uint64
	Hour       uint64
	Minute     uint64
	Second     uint64
	Nanosecond uint64
}

func (dt *DateTime) Unmarshal(b []byte) error {
	var i, j int

	dt.Hour, dt.Minute, dt.Second, dt.Nanosecond = 0, 0, 0, 0

	dt.Year, i = binary.Uvarint(b)
	if i <= 0 {
		return errors.Errorf("failed to decode datetime year (%x)", b)
	}
	dt.Month, j = binary.Uvarint(b[i:])
	if j <= 0 {
		return errors.Errorf("failed to decode datetime month (%x)", b)
	}
	i += j
	dt.Day, j = binary.Uvarint(b[i:])
	if j <= 0 {
		return errors.Errorf("failed to decode datetime day (%x)", b)
	}
	i += j
	dt.Hour, j = binary.Uvarint(b[i:])
	if j > 0 {
		i += j
		dt.Minute, j = binary.Uvarint(b[i:])
		if j > 0 {
			i += j
			dt.Second, j = binary.Uvarint(b[i:])
			if j > 0 {
				i += j
				dt.Nanosecond, j = binary.Uvarint(b[i:])
			}
		}
	}
	if j < 0 {
		return errors.Errorf("failed to decode datetime time (%x)", b)
	}
	return nil
}

func (dt *DateTime) Scan(src interface{}) error {
	switch v := src.(type) {
	case *DateTime:
		dt = v
		return nil
	}
	return errors.Errorf("unable to convert type %T to %T", src, dt)
}

type NullDateTime struct {
	DateTime
	Valid bool
}

func (dt *NullDateTime) Scan(src interface{}) error {
	if src == nil {
		dt.Valid = false
		return nil
	}
	dt.Valid = true
	switch v := src.(type) {
	case []byte:
		return dt.DateTime.Unmarshal(v)
	case DateTime:
		dt.DateTime = v
		return nil
	}
	return errors.Errorf("unable to convert type %T to %T", src, dt)
}

func parseDuration(b []byte) (time.Duration, error) {
	v, i := binary.Uvarint(b[1:])
	if i < 0 {
		return 0, errors.Errorf("failed to decode time (%x)", b)
	}
	d := time.Duration(v) * time.Hour
	if i > 0 {
		i++
		v, j := binary.Uvarint(b[i:])
		if j > 0 {
			d += time.Duration(v) * time.Minute
			i += j
			v, j = binary.Uvarint(b[i:])
			if j > 0 {
				d += time.Duration(v) * time.Second
				i += j
				v, j = binary.Uvarint(b[i:])
				d += time.Duration(v)
			}
		}
		if j < 0 {
			return 0, errors.Errorf("failed to decode time (%x)", b)
		}
	}

	if b[0] == 0x01 {
		return -d, nil
	}
	return d, nil
}
