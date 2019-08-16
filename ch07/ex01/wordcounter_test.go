package main

import "testing"

func TestWordCounter_Write(t *testing.T) {
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
			s:    "I am a cat. I have, as yet, no name.　",
			want: 10,
		},
	}
	for _, c := range cases {
		var w WordCounter
		w.Write(c.s)

		if w != WordCounter(c.want) {
			t.Fatalf("w: %d want: %d\n", w, c.want)
		}
	}
}
