package expand

import (
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	add := func(r rune) rune { return r + 1 }

	data := []struct {
		s        string
		expected string
	}{
		{
			s:        "HAL",
			expected: "",
		},
		{
			s:        "HAL$HALHAL",
			expected: "IBMIBM",
		},
		{
			s:        "HALHALHAL$",
			expected: "",
		},
	}

	for _, d := range data {
		e := expand(d.s, func(s string) string { return strings.Map(add, s) })
		if e != d.expected {
			t.Errorf("e is %s want %s\n", e, d.expected)
		}
	}
}
