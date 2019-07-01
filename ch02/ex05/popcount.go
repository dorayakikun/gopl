package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var ret int
	for x > 0 {
		x = x & (x - 1)
		ret++
	}
	return ret
}
