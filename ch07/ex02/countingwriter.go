package counting

import (
	"io"
)

type ReadCounter struct {
	io.Writer
	written int64
}

func (r *ReadCounter) Write(p []byte) (n int, err error) {
	n, err = r.Writer.Write(p)
	r.written += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	rc := ReadCounter{w, 0}
	return &rc, &rc.written
}
