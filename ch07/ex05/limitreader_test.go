package limitreader

import (
	"strings"
	"testing"
)

func TestReader_Read(t *testing.T) {
	cases := []struct {
		s    string
		n    int64
		want int
	}{
		{
			s:    "",
			n:    0,
			want: 0,
		},
		{
			s:    "a",
			n:    1,
			want: 1,
		},
		{
			s:    "abcdefg",
			n:    3,
			want: 3,
		},
		{
			s:    "あいうえお",
			n:    9,
			want: 9,
		},
	}

	for _, c := range cases {
		r := LimitReader(strings.NewReader(c.s), c.n)
		p := make([]byte, 10)

		n, _ := r.Read(p)
		if n != c.want {
			t.Fatalf("n: %d want: %d\n", n, c.want)
		}
	}
}
