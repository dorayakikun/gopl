package reverse

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	data := []struct {
		s        []byte
		expected []byte
	}{
		{
			s:        []byte("abcdefg"),
			expected: []byte("gfedcba"),
		},
		{
			s:        []byte("あいうえお"),
			expected: []byte("おえういあ"),
		},
	}
	for _, d := range data {
		actual := reverse(d.s)
		if fmt.Sprintf("%s", actual) != fmt.Sprintf("%s", d.expected) {
			t.Errorf("actual is %s, want %s", actual, d.expected)
		}
	}
}
