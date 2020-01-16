package lehmer

import "encoding/binary"

type Reader struct {
	lehmer64 uint64
}

func New(n uint64) Reader {
	return Reader{
		lehmer64: n,
	}
}

func putUint64(b []byte, v uint64) {
	_ = b[7] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
}

func (r Reader)Read(buf []byte) (n int, err error) {
	l := len(buf)
	c := l / 8
	// m := l % 8
	for i := 0; i < c; i++ {
		r.lehmer64 *= 0xda942042e4dd58b5;
		binary.LittleEndian.PutUint64(buf[i*8:i*8+8], r.lehmer64)
	}
	// r.lehmer64 *= 0xda942042e4dd58b5;
	// for i := 0; i < m; i++ {
	// 	buf[8*c+i] = byte(r.lehmer64 >> 8*uint64(i))
	// }
	return l, nil
}
