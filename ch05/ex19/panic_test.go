package panic

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	n := noop()
	fmt.Printf("n: %d\n", n)
	if n != 0 {
		t.Fatalf("n: %d want: 0\n", n)
	}
}