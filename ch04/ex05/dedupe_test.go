package dedupe

import (
	"fmt"
	"testing"
)

func TestDedupe(t *testing.T) {
	data := []struct {
		s        []string
		expected []string
	}{
		{
			s:        []string{"りんご", "りんご", "バナナ", "オレンジ", "オレンジ"},
			expected: []string{"りんご", "バナナ", "オレンジ"},
		},
		{
			s:        []string{},
			expected: []string{},
		},
		{
			s:        []string{"す", "も", "も", "も", "も", "も", "も", "も", "も", "の", "う", "ち"},
			expected: []string{"す", "も", "の", "う", "ち"},
		},
	}

	for _, d := range data {
		if fmt.Sprintf("%s", dedupe(d.s)) != fmt.Sprintf("%s", d.expected) {
			t.Errorf("s is %s, want %s", d.s, d.expected)
		}
	}
}
