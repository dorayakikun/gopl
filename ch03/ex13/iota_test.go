package iota

import (
	"testing"
)

const (
	expected1 = 1000
	expected2 = 1000000
	expected3 = 1000000000
	expected4 = 1000000000000
	expected5 = 1000000000000000
	expected6 = 1e18
	expected7 = 1e21
	expected8 = 1e24
)

func Test(t *testing.T) {
	if KB != expected1 {
		t.Errorf("%d want %d", KB, expected1)
	}
	if MB != expected2 {
		t.Errorf("%d want %d", MB, expected2)
	}
	if GB != expected3 {
		t.Errorf("%d want %d", GB, expected3)
	}
	if TB != expected4 {
		t.Errorf("%d want %d", TB, expected4)
	}
	if PB != expected5 {
		t.Errorf("%d want %d", PB, expected5)
	}
	// ここからintがoverflowするのでfloat64として受ける
	if EB != expected6 {
		t.Errorf("%g want %g", EB, expected6)
	}
	if ZB != expected7 {
		t.Errorf("%g want %g", ZB, expected7)
	}
	if YB != expected8 {
		t.Errorf("%g want %g", YB, expected8)
	}
}
