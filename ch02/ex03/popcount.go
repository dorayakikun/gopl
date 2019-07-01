package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var ret byte
	for i := 0; i < 8; i += 8 {
		ret += pc[byte(x>>(uint(i)))]
	}
	return int(ret)
}
