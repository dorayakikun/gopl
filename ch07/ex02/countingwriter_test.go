package counting

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	var buf bytes.Buffer
	w, n := CountingWriter(&buf)

	w.Write([]byte("hello"))
	w.Write([]byte("hello"))
	fmt.Printf("w: %+v n: %d want: %d\n", w, *n, 5)
	// if *n != 5 {
	// 	t.Fatalf("w: %+v n: %d want: %d\n", w, *n, 5)
	// }

	// cases := []struct {
	// 	s    string
	// 	want int64
	// }{
	// 	{
	// 		s:    "",
	// 		want: 0,
	// 	},
	// 	{
	// 		s:    "cat",
	// 		want: 3,
	// 	},
	// 	{
	// 		s:    "I am a cat. I have, as yet, no name.　",
	// 		want: 10,
	// 	},
	// }

	// for _, c := range cases {
	// 	w.Write([]byte(c.s))
	// 	if *n != c.want {
	// 		t.Fatalf("n: %d want: %d\n", *n, c.want)
	// 	}
	// }
}
