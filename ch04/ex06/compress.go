package trim

import (
	"unicode"
	"unicode/utf8"
)

func trim(bytes []byte) []byte {
	// space count
	var count int
	// 先頭空白スペースの位置
	var position int
	// 取り除かれたスペースの数
	var gap int
	// 読み込んだ文字数
	var read int
	for read < len(bytes) {
		r, size := utf8.DecodeRune(bytes[read:])
		if unicode.IsSpace(r) {
			if count == 0 {
				// 先頭のスペース位置をメモ
				position = read
			}
			read += size
			count++
			continue
		}
		// スペースが連続していた場合、２文字目以降はスペース以外の文字で詰める
		if count > 1 {
			copy(bytes[position+1:], bytes[read:])
			gap += count - 1
			count = 0
		}
		read += size
	}

	// 末尾にスペースが連続した場合の対応
	// count++で余分に足している分を含めて一つ多めに減らす(-2の部分)
	if count > 1 {
		gap += count - 2
	}
	// 取り除かれたスペース分スライスを詰める
	return bytes[:len(bytes)-gap]
}
