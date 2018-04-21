package solution

import (
	"testing"
)

type testcase struct {
	in   int
	want int
}

func TestCountPrimies(t *testing.T) {
	cases := []testcase{
		{0, 0},
		{1, 0},
		{2, 0},
		{10, 4},
		{499979, 41537},
	}
	for _, c := range cases {
		if got := countPrimes(c.in); got != c.want {
			t.Errorf("Expected '%d' but got '%d'", c.want, got)
		}
	}
}
