package iota

// 1000 ^ iota としたいが、桁溢れを起こすので(定数ではmath.Pow(x, y)が利用できない)
const (
	B  = 1
	KB = B * 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000.0
	ZB = EB * 1000.0
	YB = ZB * 1000.0
)
