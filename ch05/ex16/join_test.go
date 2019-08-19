package join

import "testing"

func TestJoin(t *testing.T) {
	tests := []struct{
		sep  string
		a    []string
		want string
	} {
		{
			sep:  ", ",
			a:    []string{ "apple", "orange", "banana" },
			want: "apple, orange, banana",
		},
		{
			sep:  "-",
			a:    []string{},
			want: "",
		},
		{
			sep:  ".",
			a:    []string{"aaa", "bbb"},
			want: "aaa.bbb",
		},
		{
			sep:  ".",
			a:    []string{"soseki"},
			want: "soseki",
		},
	}
	for _, test := range tests {
		s := join(test.sep, test.a...)
		if s != test.want {
			t.Errorf("s: %s want: %s\n", s, test.want)
		}
	}
}