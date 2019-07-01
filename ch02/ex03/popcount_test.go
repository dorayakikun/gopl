package popcount

import (
	"gopl.io/ch2/popcount"
	"testing"
)

func TestPopCount(t *testing.T) {
	ret := PopCount(19)
	if ret != 3 {
		t.Errorf("予期した値: 3、実際の値: %d", ret)
	}
}

func BenchmarkPopCountExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(uint64(i))
	}
}

func BenchmarkPopCountFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}
