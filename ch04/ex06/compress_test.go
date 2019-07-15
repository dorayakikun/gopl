package trim

import (
	"fmt"
	"testing"
)

func TestTrim(t *testing.T) {
	data := []struct {
		s        []byte
		expected []byte
	}{
		{
			s:        []byte("     Hello,  world."),
			expected: []byte(" Hello, world."),
		},
		{
			s:        []byte("Hello,               world."),
			expected: []byte("Hello, world."),
		},
		{
			s:        []byte("Hello, world.       "),
			expected: []byte("Hello, world. "),
		},
	}
	for _, d := range data {
		actual := trim(d.s)
		if fmt.Sprintf("%s", actual) != fmt.Sprintf("%s", d.expected) {
			t.Errorf("actual is %q, want %q\n", actual, d.expected)
		}
	}
}
