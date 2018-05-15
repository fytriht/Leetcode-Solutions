package solution

import (
	"testing"
)

func TestMinMoves(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 1, 1}, 0},
		{[]int{1, 2, 3}, 3},
	}

	for _, c := range cases {
		if got := minMoves(c.in); got != c.want {
			t.Errorf(
				"\n input: %v \n got: %d \n want: %d",
				c.in,
				got,
				c.want,
			)
		}
	}
}
