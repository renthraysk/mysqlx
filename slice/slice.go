package slice

func Allocate(in []byte, n int) ([]byte, []byte) {
	if nn := len(in) + n; cap(in) >= nn {
		return in[:nn], in[len(in):nn]
	}
	return in, make([]byte, n, 2*n)
}

func ForAppend(in []byte, n int) (head, tail []byte) {
	if nn := len(in) + n; cap(in) >= nn {
		head = in[:nn]
	} else {
		head = make([]byte, nn, 2*nn)
		copy(head, in)
	}
	tail = head[len(in):]
	return
}
