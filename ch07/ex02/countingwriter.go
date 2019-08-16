package counting

import (
	"io"
)

type ReadCounter struct {
	w       io.Writer
	written int64
}

func (r *ReadCounter) Write(p []byte) (n int, err error) {
	n, err = r.w.Write(p)
	r.written += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	rc := ReadCounter{w: w, written: 0}
	return &rc, &rc.written
}
