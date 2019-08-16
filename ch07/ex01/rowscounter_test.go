package main

import "testing"

func TestRowsCounter_Write(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{
			s:    "",
			want: 0,
		},
		{
			s:    "hello",
			want: 1,
		},
		{
			s:    "hello\nhi\n",
			want: 2,
		},
	}
	for _, c := range cases {
		var r RowCounter
		r.Write(c.s)

		if r != RowCounter(c.want) {
			t.Fatalf("r: %d want: %d\n", r, c.want)
		}
	}
}
