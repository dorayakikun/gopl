package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var ret int
	for i := uint64(0); i < 64; i++ {
		// 最下位ビットを比べる
		ret += int(x & 1)
		x >>= 1
	}
	return ret
}
