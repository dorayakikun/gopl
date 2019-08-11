package min

import "testing"

func TestMin(t *testing.T) {
	data := []struct {
		vals []int
		expected int
	}{
		{
			vals: []int{1, 2, 3, 4, 5, 6},
			expected: 1,
		},
		{
			vals: []int{},
			expected: 0,
		},
	}

	for _, d := range data {
		m := min(d.vals...)
		if m != d.expected {
			t.Errorf("m is %d want %d\n", m, d.expected)
		}
	}
}

func TestMin2(t *testing.T) {
	data := []struct {
		val int
		vals []int
		expected int
	}{
		{
			val: 8,
			vals: []int{1, 2, 3, 4, 5, 6},
			expected: 1,
		},
		{
			val: 0,
			vals: []int{},
			expected: 0,
		},
	}

	for _, d := range data {
		m := min2(d.val, d.vals...)
		if m != d.expected {
			t.Errorf("m is %d want %d\n", m, d.expected)
		}
	}
}