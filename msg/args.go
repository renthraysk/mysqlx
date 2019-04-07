package msg

import (
	"database/sql/driver"
	"reflect"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/renthraysk/mysqlx/collation"
)

type Args interface {
	AppendArgNull()
	AppendArgBool(bool)

	AppendArgUint(uint64)
	AppendArgInt(int64)

	AppendArgString(string, collation.Collation)
	AppendArgBytes([]byte, ContentType)

	AppendArgFloat32(float32)
	AppendArgFloat64(float64)
}

type ArgAppender interface {
	AppendArg(Args) error
}

var zeroTime = []byte{'0', '0', '0', '0', '-', '0', '0', '-', '0', '0'}

const smallsString = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859"

func appendArgValue(a Args, value interface{}) error {
derefLoop:
	if value == nil {
		a.AppendArgNull()
		return nil
	}
	switch v := value.(type) {
	case string:
		a.AppendArgString(v, 0)
	case []byte:
		a.AppendArgBytes(v, ContentTypePlain)
	case uint:
		a.AppendArgUint(uint64(v))
	case uint8:
		a.AppendArgUint(uint64(v))
	case uint16:
		a.AppendArgUint(uint64(v))
	case uint32:
		a.AppendArgUint(uint64(v))
	case uint64:
		a.AppendArgUint(v)
	case int:
		a.AppendArgInt(int64(v))
	case int8:
		a.AppendArgInt(int64(v))
	case int16:
		a.AppendArgInt(int64(v))
	case int32:
		a.AppendArgInt(int64(v))
	case int64:
		a.AppendArgInt(v)
	case bool:
		a.AppendArgBool(v)
	case float32:
		a.AppendArgFloat32(v)
	case float64:
		a.AppendArgFloat64(v)
	case time.Time:
		if v.IsZero() {
			a.AppendArgBytes(zeroTime, ContentTypePlain)
			return nil
		}

		const fmt = "2006-01-02 15:04:05.999999999"
		var b [len(fmt) + 16]byte

		a.AppendArgBytes(v.AppendFormat(b[:0], fmt), ContentTypePlain)

	case time.Duration:
		var buf [1 + 20 + 1 + 2 + 1 + 2]byte

		b := buf[:0]
		if v < 0 {
			v = -v
			b = buf[:1]
			b[0] = '-'
		}
		if v > 838*time.Hour+59*time.Minute+59*time.Second {
			return errors.New("time.Duration outside TIME range [-838:59:59, 838:59:59]")
		}

		b = strconv.AppendUint(b, uint64(v/time.Hour), 10)
		i := 2 * (uint(v/time.Minute) % 60)
		j := 2 * (uint(v/time.Second) % 60)
		b = append(b, ':', smallsString[i], smallsString[i+1], ':', smallsString[j], smallsString[j+1])
		a.AppendArgBytes(b, ContentTypePlain)

	default:
		if aa, ok := v.(ArgAppender); ok {
			return aa.AppendArg(a)
		}
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Ptr:
			if rv.IsNil() {
				a.AppendArgNull()
				return nil
			}
			value = rv.Elem().Interface()
			goto derefLoop

		default:
			return errors.Errorf("unsupported type %T, a %s", value, rv.Kind())
		}
	}
	return nil
}

func appendArgValues(a Args, values []driver.Value) error {
	for i, arg := range values {
		if err := appendArgValue(a, arg); err != nil {
			return errors.Wrapf(err, "unable to serialize argument %d", i)
		}
	}
	return nil
}

func AppendArgValues(a Args, values ...interface{}) error {
	for i, arg := range values {
		if err := appendArgValue(a, arg); err != nil {
			return errors.Wrapf(err, "unable to serialize argument %d", i)
		}
	}
	return nil
}

func appendArgNamedValues(a Args, values []driver.NamedValue) error {
	for _, arg := range values {
		if len(arg.Name) > 0 {
			return errors.New("mysql does not support the use of named parameters")
		}
		if err := appendArgValue(a, arg.Value); err != nil {
			return errors.Wrapf(err, "unable to serialize named argument %d", arg.Ordinal)
		}
	}
	return nil
}
