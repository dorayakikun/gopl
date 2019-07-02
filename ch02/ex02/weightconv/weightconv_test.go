package weightconv

import "testing"

func TestPound_String(t *testing.T) {
	p := Pound(12)
	if p.String() != "12lb" {
		t.Errorf("予期した結果: 12lb、実際の結果: %s", p.String())
	}
}

func TestKilogram_String(t *testing.T) {
	k := Kilogram(25)
	if k.String() != "25kg" {
		t.Errorf("予期した結果: 25kg、実際の結果: %s", k.String())
	}
}