package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var ret int
	for i := 0; i < 64; i++ {
		if pc[byte(x>>(uint(i)))] == 1 {
			ret++
		}
	}
	return ret
}
