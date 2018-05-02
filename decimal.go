package mysqlx

import "fmt"

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
