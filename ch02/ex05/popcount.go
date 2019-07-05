package popcount

func PopCount(x uint64) int {
	var ret int
	for x > 0 {
		x = x & (x - 1)
		ret++
	}
	return ret
}
