package comma

import "bytes"

func comma(s string) string {
	if len(s) == 3 {
		return s
	}

	var buf bytes.Buffer
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
