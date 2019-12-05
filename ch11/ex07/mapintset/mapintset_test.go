package mapintset

import (
	"math/rand"
	"testing"
	"time"
)

func TestMapIntSet_Len(t *testing.T) {
	var x MapIntSet
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

func TestMapIntSet_Remove(t *testing.T) {
	var x MapIntSet
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

func TestMapIntSet_Clear(t *testing.T) {
	var x MapIntSet
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

var seed = time.Now().UnixNano()
var out MapIntSet

func BenchmarkIntSet_Add(b *testing.B) {
	var x MapIntSet
	for i := 0; i < b.N; i++ {
		MapIntSet_Add(&x)
	}
	out = x
}
func MapIntSet_Add(x *MapIntSet) {
	rand.Seed(seed)
	x.Add(rand.Intn(1000000))
}

func BenchmarkMapIntSet_UnionWith(b *testing.B) {
	var x MapIntSet
	for i := 0; i < b.N; i++ {
		MapIntSet_UnionWith(&x)
	}
	out = x
}

func MapIntSet_UnionWith(x *MapIntSet) {
	var y MapIntSet
	rand.Seed(seed)
	y.Add(rand.Intn(1000000))
	x.UnionWith(&y)
}
