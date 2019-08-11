package join

import "testing"

func TestJoin(t *testing.T) {
	data := []struct{
		sep string
		a []string
		expexted string
	} {
		{
			sep: ", ",
			a: []string{ "apple", "orange", "banana" },
			expexted: "apple, orange, banana",
		},
		{
			sep: "-",
			a: []string{},
			expexted: "",
		},
		{
			sep: ".",
			a: []string{"aaa", "bbb"},
			expexted: "aaa.bbb",
		},
		{
			sep: ".",
			a: []string{"soseki"},
			expexted: "soseki",
		},
	}
	for _, d := range data {
		s := join(d.sep, d.a...)
		if s != d.expexted {
			t.Errorf("s is %s want %s\n", s, d.expexted)
		}
	}
}