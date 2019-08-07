package expand

import (
	"strings"
)

func expand(s string, f func(string) string) string {
	i := strings.Index(s, "$")
	if i < 0 {
		return ""
	}

	return f(string(s[i+1:]))
}
