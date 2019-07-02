package tempconv

import "testing"

func TestCelsius_String(t *testing.T) {
	c := Celsius(28)
	if c.String() != "28℃" {
		t.Errorf("予期した結果: 28℃、実際の結果: %s", c.String())
	}
}

func TestFahrenheit_String(t *testing.T) {
	f := Fahrenheit(80)
	if f.String() != "80°F" {
		t.Errorf("予期した結果: 80°F、実際の結果: %s", f.String())
	}
}

func TestKelvin_String(t *testing.T) {
	k := Kelvin(0)
	if k.String() != "0K" {
		t.Errorf("予期した結果: 0k、実際の結果: %s", k.String())
	}
}
