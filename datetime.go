package mysqlx

import (
	"encoding/binary"
	"fmt"

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

type Time struct {
	Negative   bool
	Hour       uint64
	Minute     uint64
	Second     uint64
	Nanosecond uint64
}

func (t *Time) Unmarshal(b []byte) error {
	var i, j int

	t.Hour, t.Minute, t.Second, t.Nanosecond = 0, 0, 0, 0

	t.Negative = b[0] == 0x01
	t.Hour, i = binary.Uvarint(b[1:])
	if i < 0 {
		return errors.Errorf("failed to decode time (%x)", b)
	}
	if i > 0 {
		i++
		t.Minute, j = binary.Uvarint(b[i:])
		if j > 0 {
			i += j
			t.Second, j = binary.Uvarint(b[i:])
			if j > 0 {
				i += j
				t.Nanosecond, j = binary.Uvarint(b[i:])
			}
		}
		if j < 0 {
			return errors.Errorf("failed to decode time (%x)", b)
		}
	}
	return nil
}

func (t Time) String() string {
	if t.Negative {
		if t.Nanosecond > 0 {
			return fmt.Sprintf("-%02d:%02d:%02d.%09d", t.Hour, t.Minute, t.Second, t.Nanosecond)
		}
		return fmt.Sprintf("-%02d:%02d:%02d", t.Hour, t.Minute, t.Second)
	}
	if t.Nanosecond > 0 {
		return fmt.Sprintf("%02d:%02d:%02d.%09d", t.Hour, t.Minute, t.Second, t.Nanosecond)
	}
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour, t.Minute, t.Second)
}

type NullTime struct {
	Time
	Valid bool
}

func (t *NullTime) Scan(src interface{}) error {
	if src == nil {
		t.Valid = false
		return nil
	}
	t.Valid = true
	switch v := src.(type) {
	case []byte:
		return t.Time.Unmarshal(v)
	case Time:
		t.Time = v
		return nil
	}
	return errors.Errorf("unable to convert type %T to %T", src, t)
}
