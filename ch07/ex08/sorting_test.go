package main

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

func TestCompareSort(t *testing.T) {
	start := time.Now()
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		return false
	}})
	a := fmt.Sprintf("%v", tracks)
	fmt.Printf("sort.Sort() elapsed time: %gs\n", time.Since(start).Seconds())

	start = time.Now()
	sort.Stable(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		return false
	}})
	b := fmt.Sprintf("%v", tracks)
	fmt.Printf("sort.Stable() elapsed time: %gs\n", time.Since(start).Seconds())

	if a != b {
		t.Fatalf("\na: %s\nb: %s\n", a, b)
	}
}