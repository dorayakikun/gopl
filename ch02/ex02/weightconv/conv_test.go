package weightconv

import (
	"testing"
)

func TestPToK(t *testing.T) {
	p := Pound(2.2046)
	k := PToK(p)
	if k != 1 {
		t.Errorf("予期した結果: 1、実際の結果: %g", k)
	}
}

func TestKToP(t *testing.T) {
	k := Kilogram(1)
	p := KToP(k)
	if p != 2.2046 {
		t.Errorf("予期した結果: 2.2046、実際の結果: %g", p)
	}
}
