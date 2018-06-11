package solution

import (
	"testing"
)

func TestDistributeCandies(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 1, 2, 2, 3, 3}, 3},
		{[]int{1, 1, 2, 3}, 2},
		{[]int{1, 1, 1, 1}, 1},
		{[]int{1, 1, 2, 2}, 2},
	}

	for _, c := range cases {
		if got := distributeCandies(c.in); got != c.want {
			t.Errorf(
				"\n input: %v \n want: %d \n got: %d",
				c.in,
				c.want,
				got,
			)
		}
	}
}
