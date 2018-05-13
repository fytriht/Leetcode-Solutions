package solution

import (
	"testing"
)

func TestArrangeCoins(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 2},
		{5, 2},
		{6, 3},
		{7, 3},
		{8, 3},
		{9, 3},
		{10, 4},
	}

	for _, c := range cases {
		if got := arrangeCoins(c.in); got != c.want {
			t.Errorf(
				"\n in: %d \n got: %d \n want: %d",
				c.in,
				got,
				c.want,
			)
		}
	}
}
