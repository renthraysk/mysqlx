package mysqlx

import (
	"math/big"

	"github.com/cockroachdb/apd"
	"github.com/pkg/errors"
)

type Decimal struct {
	apd.Decimal
}

var (
	ten     = big.NewInt(10)
	hundred = big.NewInt(100)
)

func (d *Decimal) Unmarshal(b []byte) error {
	var y big.Int
	var z uint8

	d.Form = apd.Finite
	d.Exponent = -int32(b[0])
	x := &d.Coeff
	for _, z = range b[1:] {
		if z > 0x9F {
			break
		}
		if z&0x0F > 0x09 {
			x.Mul(x, ten)
			x.Add(x, y.SetUint64(uint64(z>>4)))
			break
		}
		x.Mul(x, hundred)
		x.Add(x, y.SetUint64(uint64((z>>4)*10+(z&0x0F))))
	}
	d.Negative = z == 0xD0 || z&0x0F == 0x0D
	return nil
}

type NullDecimal struct {
	Decimal
	Valid bool
}

func (nd *NullDecimal) Scan(src interface{}) error {
	if src == nil {
		nd.Valid = false
		return nil
	}
	nd.Valid = true
	switch v := src.(type) {
	case []byte:
		return nd.Decimal.Unmarshal(v)
	case Decimal:
		nd.Decimal = v
		return nil
	}
	return errors.Errorf("unable to convert %T to Decimal", src)
}
