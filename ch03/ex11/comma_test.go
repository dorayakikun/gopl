package comma

import "testing"

const (
	expected1 = "123"
	expected2 = "12,345"
	expected3 = "1,234,567"
	expected4 = "123.456"
	expected5 = "+12,345,67"
	expected6 = "-1,234,567.89"
)

func TestComma(t *testing.T) {
	s := comma("123")
	if s != expected1 {
		t.Errorf("s is %q want: %q", s, expected1)
	}

	s = comma("12345")
	if s != expected2 {
		t.Errorf("s is %q want: %q", s, expected2)
	}

	s = comma("1234567")
	if s != expected3 {
		t.Errorf("s is %q want: %q", s, expected3)
	}

	s = comma("123.456")
	if s != expected1 {
		t.Errorf("s is %q want: %q", s, expected1)
	}

	s = comma("+12345.67")
	if s != expected2 {
		t.Errorf("s is %q want: %q", s, expected2)
	}

	s = comma("-1234567.89")
	if s != expected3 {
		t.Errorf("s is %q want: %q", s, expected3)
	}
}
