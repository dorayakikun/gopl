package max

func max(vals ...int) int {
	max := 0
	// maxのゼロ値が判定対象とならないように先頭の値を先に格納する
	if len(vals) > 0 {
		max = vals[0]
		vals = vals[1:]
	}

	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

func max2(val int, vals ...int) int {
	max:= val
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}