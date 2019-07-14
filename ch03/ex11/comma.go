package comma

import (
	"bytes"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer
	var prefix string
	var suffix string
	if strings.HasPrefix(s, "+") {
		prefix = s[:1]
		s = s[1:]
	}
	if strings.HasPrefix(s, "-") {
		prefix = s[:1]
		s = s[1:]
	}
	dot := strings.LastIndex(s, ".")
	if dot != -1 {
		suffix = s[dot:]
		s = s[:dot]
	}
	if len(s) == 3 {
		buf.WriteString(prefix)
		buf.WriteString(s)
		buf.WriteString(suffix)
		return s
	}

	// カンマより左側にはみ出ている文字数
	c := len(s) % 3

	// カンマより左側の数値を書き出す
	buf.WriteString(s[:c])

	s = s[c:]
	for len(s) >= 3 {
		buf.WriteString(",")
		buf.WriteString(s[:3])
		s = s[3:]
	}
	return buf.String()
}
