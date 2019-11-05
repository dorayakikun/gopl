package popcount

import "testing"

var input uint64 = 0x1234567890ABCDEF

var out int

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = PopCount(input)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = PopCount2(input)
	}
}