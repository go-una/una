package tools

import (
	"testing"
)

func TestRand(t *testing.T) {
	for i := 0; i < 20; i++ {
		r := Rand(0, 5)
		if r < 0 || r > 5 {
			t.Fatalf("rand range expect in [0, 5], actual: %d", r)
		}
	}

	r2 := Rand(2, 2)
	if r2 != 2 {
		t.Fatalf("rand result expect %d, actual %d", 2, r2)
	}
}
