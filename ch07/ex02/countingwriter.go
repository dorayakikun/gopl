package counting

import (
	"io"
)

type MyWriter struct {
	w       io.Writer
	written int64
}

func (m MyWriter) Write(p []byte) (n int, err error) {
	n, err = m.w.Write(p)
	m.written += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	m := MyWriter{w: w, written: 0}
	return m, &m.written
}
