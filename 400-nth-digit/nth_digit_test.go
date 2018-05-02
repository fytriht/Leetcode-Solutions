package solution

import (
	"testing"
)

func TestFindNthDigit(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{1, 1},
		{3, 3},
		{11, 0},
		{12, 1},
		{1000, 3},
	}
	for _, c := range cases {
		if got := findNthDigit(c.in); got != c.want {
			t.Errorf(
				"\n input: %d \n got: %d \n want: %d",
				c.in,
				got,
				c.want,
			)
		}
	}
}

// 1 2 3 4 5 6 7 8 9 1
// 0 1 1 1 2 1 3 1 4 1
// 5 1 6 1 7 1 8 1 9 2
// 0 2 1 2 2 2 3 2 4 2
