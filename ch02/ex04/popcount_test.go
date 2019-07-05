package popcount

import (
	"gopl.io/ch2/popcount"
	"strconv"
	"testing"
)

func TestPopCount(t *testing.T) {
	ret := PopCount(10)
	if ret != 2 {
		t.Errorf("予期した値: 2、実際の値: %d", ret)
	}
}

var input, _ = strconv.ParseUint("11111", 2, 0)
var output int

func BenchmarkPopCountExpression(b *testing.B) {
	var temp int
	for i := 0; i < b.N; i++ {
		temp += popcount.PopCount(input)
	}
	output = temp
}

func BenchmarkPopCount(b *testing.B) {
	var temp int
	for i := 0; i < b.N; i++ {
		temp += PopCount(input)
	}
	output = temp
}
