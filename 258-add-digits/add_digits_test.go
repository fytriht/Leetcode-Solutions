package solution

import (
	"testing"
)

func TestAddDigits(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{0, 0},
		{1, 1},
		{38, 2},
		{948, 3},
	}

	for _, c := range cases {
		if got := addDigits(c.in); got != c.want {
			t.Errorf(
				"addDigits(%d): expected %d but got %d",
				c.in,
				c.want,
				got,
			)
		}
	}
}
