package cyclic

import "testing"

func TestIsCyclic(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	type parent struct {
		*person
		child *person
	}

	alice := &person{"Alice", 17}
	ruby := &person{"Ruby", 28}

	cases := []struct {
		input interface{}
		want  bool
	}{
		{
			0,
			false,
		},
		{
			parent{alice, alice},
			true,
		},
		{
				parent{ruby, alice},
				false,
		},
		{
			[]*person{alice, alice},
			true,
		},
		{
			[]*person{alice, ruby},
			false,
		},
	}

	for _, c := range cases {
		if ret := IsCyclic(c.input); ret != c.want {
			t.Errorf("isCyclic(%v) failed, want %t\n", c.input, ret)
		}
	}
}
