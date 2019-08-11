package join

import "strings"

func join(sep string, a ...string) string {
	var builder strings.Builder

	if len(a) == 0 {
		return ""
	}

	builder.WriteString(a[0])
	a = a[1:]
	for _, s := range a {
		builder.WriteString(sep)
		builder.WriteString(s)
	}

	return builder.String()
}
