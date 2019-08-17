package palindrome

import "sort"

func isPalindrome(s sort.Interface) bool {
	half := s.Len() / 2
	j := s.Len()-1
	for i := 0; i < half; i++ {
		if s.Less(i, j) || s.Less(i, j) {
			return false
		}
	}
	return true
}
