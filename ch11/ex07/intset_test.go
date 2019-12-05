package intset

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestIntSet_Len(t *testing.T) {
	var x IntSet
	if x.Len() != 0 {
		t.Errorf("x.Len() is %d want %d", x.Len(), 0)
		return
	}

	x.Add(1)
	if x.Len() != 1 {
		t.Errorf("x.Len() is %d want %d", x.Len(), 1)
		return
	}

	x.Add(63)
	if x.Len() != 2 {
		t.Errorf("x.Len() is %d want %d", x.Len(), 1)
		return
	}

	// 同じ場所なので、長さは変わらない
	x.Add(1)
	if x.Len() != 2 {
		t.Errorf("x.Len() is %d want %d", x.Len(), 2)
		return
	}

	x.Add(2)
	if x.Len() != 3 {
		t.Errorf("x.Len() is %d want %d", x.Len(), 3)
		return
	}
}

func TestIntSet_Remove(t *testing.T) {
	var x IntSet
	vals := []int{5, 128}

	for _, val := range vals {
		x.Add(val)
	}

	for _, val := range vals {
		x.Remove(val)
		if x.Has(val) {
			t.Fatalf("remove failed: %+v", x.words)
		}
	}
}

func TestIntSet_Clear(t *testing.T) {
	var x IntSet
	x.Add(100)

	if x.Len() != 1 {
		t.Errorf("x.Len() is %d want %d", x.Len(), 1)
		return
	}

	x.Clear()
	if x.Len() != 0 {
		t.Errorf("x.Len() is %d want %d", x.Len(), 0)
		return
	}
}

func TestIntSet_Copy(t *testing.T) {
	var x IntSet
	x.Add(100)

	y := x.Copy()

	if x.String() != y.String() {
		t.Errorf("x is %s but y is %s", x.String(), y.String())
	}
}

func TestIntSet_AddAll(t *testing.T) {
	var x IntSet
	vals := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75}

	x.AddAll(vals...)
	for _, val := range vals {
		if !x.Has(val) {
			t.Fatalf("add %d failed: %s\n", val, x.String())
		}
	}
}

func TestIntSet_IntersectWith(t *testing.T) {
	cases := []struct {
		xs   []int
		ys   []int
		want []int
	}{
		{
			xs:   []int{1, 2, 3, 4, 5},
			ys:   []int{3, 4, 5, 6, 7, 8},
			want: []int{3, 4, 5},
		},
		{
			xs:   []int{63, 62, 61, 60},
			ys:   []int{63, 64, 65, 66},
			want: []int{63},
		},
		{
			xs:   []int{70, 71, 72, 73},
			ys:   []int{1, 2, 3},
			want: []int{70, 71, 72, 73},
		},
	}

	for _, c := range cases {
		var x IntSet
		x.AddAll(c.xs...)
		var y IntSet
		y.AddAll(c.ys...)

		x.IntersectWith(&y)
		for _, w := range c.want {
			if !x.Has(w) {
				t.Fatalf("missing %d: %s\n", w, x.String())
			}
		}
	}
}

func TestIntSet_DifferenceWith(t *testing.T) {
	cases := []struct {
		xs   []int
		ys   []int
		want []int
	}{
		{
			xs:   []int{1, 2, 3, 4, 5},
			ys:   []int{3, 4, 5, 6, 7, 8},
			want: []int{1, 2},
		},
		{
			xs:   []int{63, 62, 61, 60},
			ys:   []int{63, 64, 65, 66},
			want: []int{60, 61, 62},
		},
		{
			xs:   []int{70, 71, 72, 73},
			ys:   []int{1, 2, 3},
			want: []int{70, 71, 72, 73},
		},
	}

	for _, c := range cases {
		var x IntSet
		x.AddAll(c.xs...)
		var y IntSet
		y.AddAll(c.ys...)

		x.DifferenceWith(&y)
		for _, w := range c.want {
			if !x.Has(w) {
				t.Fatalf("missing %d: %s\n", w, x.String())
			}
		}
	}
}

func TestIntSet_SymmetricDifference(t *testing.T) {
	cases := []struct {
		xs   []int
		ys   []int
		want []int
	}{
		{
			xs:   []int{1, 2, 3, 4, 5},
			ys:   []int{3, 4, 5, 6, 7, 8},
			want: []int{1, 2, 6, 7, 8},
		},
		{
			xs:   []int{63, 62, 61, 60},
			ys:   []int{63, 64, 65, 66},
			want: []int{60, 61, 62, 64, 65, 66},
		},
		{
			xs:   []int{70, 71, 72, 73},
			ys:   []int{1, 2, 3},
			want: []int{1, 2, 3, 70, 71, 72, 73},
		},
	}

	for _, c := range cases {
		var x IntSet
		x.AddAll(c.xs...)
		var y IntSet
		y.AddAll(c.ys...)

		x.SymmetricDifference(&y)
		for _, w := range c.want {
			if !x.Has(w) {
				t.Fatalf("missing %d: %s\n", w, x.String())
			}
		}
	}
}

func TestIntSet_Elems(t *testing.T) {
	cases := []struct {
		xs   []int
		want []int
	}{
		{
			xs:   []int{1, 2, 3, 4, 5, 64, 65, 66, 129, 130},
			want: []int{1, 2, 3, 4, 5, 64, 65, 66, 129, 130},
		},
		{
			xs:   []int(nil),
			want: []int(nil),
		},
		{
			xs:   []int{1, 1, 1, 1, 1, 1, 1, 1, 1},
			want: []int{1},
		},
	}

	for _, c := range cases {
		var x IntSet
		x.AddAll(c.xs...)

		elems := x.Elems()

		if len(elems) != len(c.want) {
			t.Fatalf("elems: %+v want %+v\n", elems, c.want)
		}

		for i, e := range elems {
			if e != c.want[i] {
				t.Fatalf("elems: %+v want %+v\n", elems, c.want)
			}
		}
	}
}

var seed = time.Now().UnixNano()
var out IntSet

func BenchmarkIntSet_Add(b *testing.B) {
	var x IntSet
	for i := 0; i < b.N; i++ {
		IntSet_Add(&x)
	}
	out = x
}
func IntSet_Add(x *IntSet) {
	rand.Seed(seed)
	x.Add(rand.Intn(1000000))
}

func BenchmarkIntSet_UnionWith(b *testing.B) {
	var x IntSet
	for i := 0; i < b.N; i++ {
		IntSet_UnionWith(&x)
	}
	out = x
}

func IntSet_UnionWith(x *IntSet) {
	var y IntSet
	rand.Seed(seed)
	y.Add(rand.Intn(1000000))
	x.UnionWith(&y)
}
