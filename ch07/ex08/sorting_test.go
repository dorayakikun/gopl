package main

import (
	"sort"
	"testing"
)

// TODO customSortと sort.stable x 2をbenchで比較

type customSort2 struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort2) Len() int           { return len(x.t) }
func (x customSort2) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort2) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Sort(customSort2{tracks, func(x, y *Track) bool {
			if x.Artist != y.Artist {
				return x.Artist < y.Artist
			}
			if x.Title != y.Title {
				return x.Title < y.Title
			}
			return false
		}})
	}
}

func BenchmarkStable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Stable(customSort2{tracks, func(x, y *Track) bool {
			if x.Artist != y.Artist {
				return x.Artist < y.Artist
			}
			if x.Title != y.Title {
				return x.Title < y.Title
			}
			return false
		}})
	}
}