package lengthconv

import "testing"

func TestFeet_String(t *testing.T) {
	f := Feet(10.5)
	if f.String() != "10.5ft" {
		t.Errorf("予期した結果: 10.5ft、実際の結果: %s", f.String())
	}
}

func TestMeters_String(t *testing.T) {
	m := Meters(10.5)
	if m.String() != "10.5m" {
		t.Errorf("予期した結果: 10.5m、実際の結果: %s", m.String())
	}
}