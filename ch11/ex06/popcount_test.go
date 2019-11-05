package popcount

import "testing"

var out int
var out2 int

func BenchmarkPopCount(b *testing.B) {
	var ret int
	for i := 0; i < b.N; i++ {
		ret += PopCount(0x1234567890ABCDEF)
	}
	out = ret
}

func BenchmarkPopCount2(b *testing.B) {
	var ret int
	for i := 0; i < b.N; i++ {
		ret += PopCount2(0x1234567890ABCDEF)
	}
	out2 = ret
}
