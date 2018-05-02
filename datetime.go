package mysqlx

import (
	"encoding/binary"

	"github.com/pkg/errors"
)

type Date struct {
	Year  uint64
	Month uint64
	Day   uint64
}

func (d *Date) Reset() {
	d.Year, d.Month, d.Day = 0, 0, 0
}

func (d *Date) Unmarshal(b []byte) error {
	var i, j int

	d.Reset()

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

type DateTime struct {
	Date
	Hour       uint64
	Minute     uint64
	Second     uint64
	Nanosecond uint64
}

func (dt *DateTime) Reset() {
	dt.Date.Reset()
	dt.Hour, dt.Minute, dt.Second, dt.Nanosecond = 0, 0, 0, 0
}

func (dt *DateTime) Unmarshal(b []byte) error {
	var i, j int

	dt.Reset()

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

type Time struct {
	Negative   bool
	Hour       uint64
	Minute     uint64
	Second     uint64
	Nanosecond uint64
}

func (t *Time) Reset() {
	t.Negative = false
	t.Hour, t.Minute, t.Second, t.Nanosecond = 0, 0, 0, 0
}

func (t *Time) Unmarshal(b []byte) error {
	var i, j int

	t.Reset()

	t.Negative = b[0] == 0x01
	t.Hour, i = binary.Uvarint(b[1:])
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
	}
	if j < 0 {
		return errors.Errorf("failed to decode time (%x)", b)
	}
	return nil
}
