package max

import "testing"

func TestMax(t *testing.T) {
	data := []struct {
		vals     []int
		expected int
	}{
		{
			vals:     []int{0, 1, 2, 3, 4, 5, 6},
			expected: 6,
		},
		{
			vals:     []int{},
			expected: 0,
		},
		{
			vals:     []int{-1, -2, -3, -4, -5, -6},
			expected: -1,
		},
	}

	for _, d := range data {
		m := max(d.vals...)
		if m != d.expected {
			t.Errorf("m is %d want %d\n", m, d.expected)
		}
	}
}

func TestMax2(t *testing.T) {
	data := []struct {
		val      int
		vals     []int
		expected int
	}{
		{
			val:      1,
			vals:     []int{0, 1, 2, 3, 4, 5, 6},
			expected: 6,
		},
		{
			val:      0,
			vals:     []int{},
			expected: 0,
		},
		{
			val:      -8,
			vals:     []int{-1, -2, -3, -4, -5, -6},
			expected: -1,
		},
	}

	for _, d := range data {
		m := max2(d.val, d.vals...)
		if m != d.expected {
			t.Errorf("m is %d want %d\n", m, d.expected)
		}
	}
}
