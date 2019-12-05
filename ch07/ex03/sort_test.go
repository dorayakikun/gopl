package treesort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestTree_String(t *testing.T) {
	cases := []struct {
		data []int
		want string
	}{
		{
			data: []int(nil),
			want: "",
		},
		{
			data: []int{0},
			want: "0\n",
		},
		{
			data: []int{0, 1, 2},
			/**
					2
				1
			0
			*/
			want: "\t\t2\n\t1\n0\n",
		},
		{
			data: []int{0, -1, -2},
			/**
			0
				-1
					-2
			*/
			want: "0\n\t-1\n\t\t-2\n",
		},
		{
			data: []int{0, -1, -2, 2, 3, 6, 5, 7},
			/**
							7
						6
							5
					3
				2
			0
				-1
					-2
			*/
			want: "\t\t\t\t7\n\t\t\t6\n\t\t\t\t5\n\t\t3\n\t2\n0\n\t-1\n\t\t-2\n",
		},
	}

	for _, c := range cases {
		var root *tree
		for _, d := range c.data {
			root = add(root, d)
		}
		s := root.String()
		if s != c.want {
			t.Fatalf("s:\n%s\n\nwant:\n%s\n", s, c.want)
		}
	}
}
