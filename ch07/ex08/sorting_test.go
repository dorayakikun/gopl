package main

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

func TestCompareSort(t *testing.T) {
	start := time.Now()
	sort.Sort(byArtist(tracks))
	sort.Sort(byYear(tracks))
	a := fmt.Sprintf("%v", tracks)
	fmt.Printf("elapsed time: %gs\n", time.Since(start).Seconds())

	start = time.Now()
	sort.Stable(byArtist(tracks))
	sort.Stable(byYear(tracks))
	b := fmt.Sprintf("%v", tracks)
	fmt.Printf("elapsed time: %gs\n", time.Since(start).Seconds())

	if a != b {
		t.Fatalf("\na: %s\nb: %s\n", a, b)
	}
}