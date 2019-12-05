package split

import (
	"fmt"
	"strings"
	"testing"
)

func assertEqual(x, y int) {
	if x != y {
		panic(fmt.Sprintf("%d != %d", x, y))
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		s    string
		sep  string
		want int
	}{
		{
			"a:b:c",
			":",
			3,
		},
		{
			"い、ろ、は、に、ほ、へ、と",
			"、",
			7,
		},
		{
			"xyz",
			",",
			1,
		},
	}

	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.s, test.sep, got, test.want)
		}
	}
}
