package reverse

import (
	"unicode/utf8"
)

func _reverse(bytes []byte) {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}

func reverse(bytes []byte) []byte {
	var offset int
	for offset < len(bytes) {
		_, size := utf8.DecodeRune(bytes)

		// 前回のループで並び替えた位置より手前のsliceを逆順sort
		_reverse(bytes[:len(bytes)-offset])

		// validなruneに並び替え
		offset += size
		position := len(bytes) - offset
		_reverse(bytes[position : position+size])

		// 対象以外の[]byteを元の順序に戻す
		_reverse(bytes[:len(bytes)-offset])
	}
	return bytes
}
