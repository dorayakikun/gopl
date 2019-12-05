package min

func min(vals ...int) int {
	min := 0

	// minのゼロ値が判定対象とならないように先頭の値を先に格納する
	if len(vals) > 0 {
		min = vals[0]
		vals = vals[1:]
	}

	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}

func min2(val int, vals ...int) int {
	min := val
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}
