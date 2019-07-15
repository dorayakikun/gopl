package dedupe

func dedupe(strings []string) []string {
	if len(strings) == 0 {
		return strings
	}

	out := strings[:0]
	var duplicate string
	for _, s := range strings {
		if s != duplicate {
			out = append(out, s)
			duplicate = s
		}
	}
	return out[:]
}
