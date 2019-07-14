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
	// ここからintがoverflowするので、溢れないように除算した結果と比較する
	if EB/1000 != expected5 {
		t.Errorf("EB/1000 want %d", expected5)
	}
	if ZB/1000000 != expected5 {
		t.Errorf("ZB/1000000 want %d", expected5)
	}
	if YB/1000000000 != expected5 {
		t.Errorf("YB/1000000 want %d", expected5)
	}
}
