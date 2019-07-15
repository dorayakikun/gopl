package rotate

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	expected := []int{5, 6, 7, 1, 2, 3, 4}
	s := []int{1, 2, 3, 4, 5, 6, 7}
	s = rotate(s, 4)

	if fmt.Sprintf("%d", s) != fmt.Sprintf("%d", expected) {
		t.Errorf("s is %d, want %d", s, expected)
	}
}
