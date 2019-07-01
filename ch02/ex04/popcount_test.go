package popcount

import (
	"gopl.io/ch2/popcount"
	"testing"
)

func TestPopCount(t *testing.T) {
	ret := PopCount(10)
	if ret != 2 {
		t.Errorf("予期した値: 2、実際の値: %d", ret)
	}
}

func BenchmarkPopCountExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(uint64(i))
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}
