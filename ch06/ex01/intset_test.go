package intset

import (
	"fmt"
	"testing"
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
	x.Add(5)
	x.Add(128)
	x.Remove(128)

	if x.String() != "{5}" {
		t.Errorf("x.String() is %s want {5}", x.String())
	}

	x.Remove(5)
	if x.String() != "{}" {
		t.Errorf("x.String() is %s want {}", x.String())
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