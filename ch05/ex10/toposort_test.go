package main

import "testing"

func TestTopoSort(t *testing.T) {
	// 自分より前に自分の依存関係

	order := topoSort(prereqs)

	for i, o := range order {
		m := prereqs[o]
		for k := range m {
			found := false
			for j := 0; j < i; j++ {
				if order[j] == k {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("missing dependencies: %q -> %q", o, k)
			}
		}
	}
}
