package anagram

import "sort"

func anagram(a string, b string) bool {
	// そもそも文字数が異なる場合は、以降の処理を実行しない
	if len(a) != len(b) {
		return false
	}

	// ルーンのスライスをソート順を揃えて、愚直にマッチング
	ra := []rune(a)
	rb := []rune(b)
	sort.Slice(ra, func(i, j int) bool {
		return ra[i] < ra[j]
	})
	sort.Slice(rb, func(i, j int) bool {
		return rb[i] < rb[j]
	})
	for i, r := range ra {
		if r != rb[i] {
			return false
		}
	}
	return true
}
