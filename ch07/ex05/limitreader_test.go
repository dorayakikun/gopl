package limitreader

import (
	"io"
	"strings"
	"testing"
)

func TestReader_Read(t *testing.T) {
	cases := []struct {
		s string
		n int64
		want int64
	} {
		{
			s: "",
			n: 0,
			want: 0,
		},
		{
			s: "a",
			n: 1,
			want: 0,
		},
		{
			s: "abcdefg",
			n: 3,
			want: 0,
		},
		{
			s: "あいうえお",
			n: 200,
			want: 185,
		},
	}

	for _, c := range cases {
		r, n := LimitReader(strings.NewReader(c.s), c.n)
		p := make([]byte, 100)
		for {
			_, err :=r.Read(p)
			if err == io.EOF {
				break
			}
		}

		if *n != c.want {
			t.Fatalf("s: %s n: %d want: %d\n", c.s, n, c.want)
		}
	}
}