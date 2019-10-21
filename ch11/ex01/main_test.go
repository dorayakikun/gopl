package main

import (
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestRun(t *testing.T) {
	tests := []struct {
		in      string
		counts  map[rune]int
		utflen  [utf8.UTFMax + 1]int
		invalid int
	}{
		{
			"",
			map[rune]int{},
			[utf8.UTFMax + 1]int{0, 0, 0, 0, 0},
			0,
		},
		{
			"apple",
			map[rune]int{'a': 1, 'p': 2, 'l': 1, 'e': 1},
			[utf8.UTFMax + 1]int{0, 5, 0, 0, 0},
			0,
		},
		{
			"apple banana",
			map[rune]int{'a': 4, 'b': 1, 'n': 2, 'p': 2, 'l': 1, 'e': 1, ' ': 1},
			[utf8.UTFMax + 1]int{0, 12, 0, 0, 0},
			0,
		},

		{
			"apple banana \xa0\xa1",
			map[rune]int{'a': 4, 'b': 1, 'n': 2, 'p': 2, 'l': 1, 'e': 1, ' ': 2},
			[utf8.UTFMax + 1]int{0, 13, 0, 0, 0},
			2,
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.in)
		counts, utflen, invalid := run(r)

		if !reflect.DeepEqual(counts, test.counts) {
			t.Errorf("counts: %v want: %v", counts, test.counts)
		}

		if !reflect.DeepEqual(utflen, test.utflen) {
			t.Errorf("utflen: %v want: %v", utflen, test.utflen)
		}

		if invalid != test.invalid {
			t.Errorf("invalid: %v want: %v", invalid, test.invalid)
		}
	}
}
