package anagram

import "testing"

func TestAnagram(t *testing.T) {
	if !anagram("elvis", "lives") {
		t.Error("false want true\n")
	}
	if anagram("すりえふ", "ふぇりす") {
		t.Error("true want false\n")
	}
	if anagram("ごー", "らすと") {
		t.Error("true want false\n")
	}
}
