package limitreader

import "io"

type reader struct {
	r io.Reader
	n int64
}

func (r *reader) Read(b []byte) (n int, err error) {
	if r.n <= 0 {
		err = io.EOF
		return
	}
	if int64(len(b)) > r.n {
		b = b[:r.n]
	}

	n, err = r.r.Read(b)
	r.n -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &reader{r, n}
}
