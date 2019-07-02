package lengthconv

import "testing"

func TestFToM(t *testing.T) {
	f := Feet(3.281)
	m := FToM(f)

	if m != 1 {
		t.Errorf("予期した結果: 1、実際の結果: %g", m)
	}
}

func TestMToF(t *testing.T) {
	m := Meters(1)
	f := MToF(m)

	if f != 3.281 {
		t.Errorf("予期した結果: 3.281、実際の結果: %g", m)
	}
}
