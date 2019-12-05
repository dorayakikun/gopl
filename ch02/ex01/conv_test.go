package tempconv

import (
	"math"
	"testing"
)

func TestCToF(t *testing.T) {
	f := CToF(0)
	if f != 32 {
		t.Errorf("予期した結果: 32、実際の結果: %.2f", f)
	}
}

func TestFtoC(t *testing.T) {
	c := FtoC(0)
	if math.Ceil(float64(c)) != -17 {
		t.Errorf("予期した結果: 17、実際の結果: %.2f", math.Ceil(float64(c)))
	}
}

func TestCToK(t *testing.T) {
	k := CToK(-273.15)
	if k != 0 {
		t.Errorf("予期した結果: 0、実際の結果: %.2f", k)
	}
}

func TestKToC(t *testing.T) {
	c := KToC(1)
	if c != -272.15 {
		t.Errorf("予期した結果: -272.15、実際の結果: %.2f", c)
	}
}
