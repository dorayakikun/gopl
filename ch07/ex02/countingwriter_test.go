package counting

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	var buf bytes.Buffer
	w, n := CountingWriter(&buf)
	cases := []struct {
		s    string
		want int64
	}{
		{
			s:    "",
			want: 0,
		},
		{
			s:    "cat",
			want: 3,
		},
		{
			s:    "I am a cat. I have, as yet, no name.　",
			want: 42,
		},
	}

	for _, c := range cases {
		w.Write([]byte(c.s))
		if *n != c.want {
			t.Fatalf("n: %d want: %d\n", *n, c.want)
		}
	}
}
