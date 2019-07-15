package rotate

func rotate(s []int, i int) []int {
	copy(s, append(s[i:], s[:i]...))
	return s[:]
}
