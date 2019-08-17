package limitreader

import "io"

type Reader struct {
	R io.Reader
	N int64
}

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.N <= 0 {
		err = io.EOF
		return
	}
	if int64(len(b)) > r.N {
		b = b[0:r.N]
	}
	n, err = r.R.Read(b)
	if err != nil {
		return
	}
	r.N -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) (io.Reader, *int64) {
	l := &Reader{r, n}
	return l, &l.N
}