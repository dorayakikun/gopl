package reverse

import "testing"

func TestReverse(t *testing.T) {
	s := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(&s)

	expected := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	if s != expected {
		t.Errorf("s is %d, want %d", s, expected)
	}
}
