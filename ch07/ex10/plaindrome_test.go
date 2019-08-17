package palindrome

import (
	"sort"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		a    []int
		want bool
	}{
		{
			a:    []int(nil),
			want: true,
		},
		{
			a:    []int{1, 2, 3},
			want: false,
		},
		{
			a:    []int{3, 2, 3},
			want: true,
		},
	}

	for _, c := range cases {
		actual := isPalindrome(sort.IntSlice(c.a))
		if actual != c.want {
			t.Fatalf("actual: %t want: %t a: %+v\n", actual, c.want, c.a)
		}
	}
}

type text []rune

func (t text) Len() int { return len(t) }
func (t text) Less(i, j int) bool { return t[i] < t[j] }
func (t text) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

func TestIsPalindrome2(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{
			s:    "",
			want: true,
		},
		{
			s:    "いろは",
			want: false,
		},
		{
			s:    "とまと",
			want: true,
		},
		{
			s:    "coc",
			want: true,
		},
	}

	for _, c := range cases {
		actual := isPalindrome(text([]rune(c.s)))
		if actual != c.want {
			t.Fatalf("actual: %t want: %t s: %s\n", actual, c.want, c.s)
		}
	}
}