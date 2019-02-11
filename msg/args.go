package msg

import (
	"fmt"
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

// AppendArgTime appends a time parameter
func appendArgTime(s Args, t time.Time) {
	const fmt = "2006-01-02 15:04:05.999999999"
	var b [len(fmt) + 16]byte

	if t.IsZero() {
		s.AppendArgBytes(zeroTime, ContentTypePlain)
		return
	}

	s.AppendArgBytes(t.AppendFormat(b[:0], fmt), ContentTypePlain)
}

const smallsString = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859"

func appendArgDuration(s Args, d time.Duration) error {
	var buf [1 + 20 + 1 + 2 + 1 + 2]byte

	b := buf[:0]
	if d < 0 {
		d = -d
		b = buf[:1]
		b[0] = '-'
	}
	if d > 838*time.Hour+59*time.Minute+59*time.Second {
		return errors.New("time.Duration outside TIME range [-838:59:59, 838:59:59]")
	}

	b = strconv.AppendUint(b, uint64(d/time.Hour), 10)
	i := 2 * (uint(d/time.Minute) % 60)
	j := 2 * (uint(d/time.Second) % 60)
	b = append(b, ':', smallsString[i], smallsString[i+1],
		':', smallsString[j], smallsString[j+1])
	s.AppendArgBytes(b, ContentTypePlain)
	return nil
}

func appendArgValue(s Args, value interface{}) error {
	if value == nil {
		s.AppendArgNull()
		return nil
	}
	switch v := value.(type) {
	case string:
		s.AppendArgString(v, 0)
	case []byte:
		s.AppendArgBytes(v, ContentTypePlain)
	case uint:
		s.AppendArgUint(uint64(v))
	case uint8:
		s.AppendArgUint(uint64(v))
	case uint16:
		s.AppendArgUint(uint64(v))
	case uint32:
		s.AppendArgUint(uint64(v))
	case uint64:
		s.AppendArgUint(v)
	case int:
		s.AppendArgInt(int64(v))
	case int8:
		s.AppendArgInt(int64(v))
	case int16:
		s.AppendArgInt(int64(v))
	case int32:
		s.AppendArgInt(int64(v))
	case int64:
		s.AppendArgInt(v)
	case bool:
		s.AppendArgBool(v)
	case float32:
		s.AppendArgFloat32(v)
	case float64:
		s.AppendArgFloat64(v)
	case time.Time:
		appendArgTime(s, v)
	case time.Duration:
		appendArgDuration(s, v)
	default:
		if a, ok := v.(ArgAppender); ok {
			return a.AppendArg(s)
		}
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Ptr:
			if rv.IsNil() {
				s.AppendArgNull()
				return nil
			}
			// @TODO RECURSIVE
			return appendArgValue(s, rv.Elem().Interface())
		default:
			return fmt.Errorf("unsupported type %T, a %s", value, rv.Kind())
		}
	}
	return nil
}
