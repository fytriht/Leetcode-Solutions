package solution

import (
	"testing"
)

func TestIslandPerimeter(t *testing.T) {
	cases := []struct {
		in   [][]int
		want int
	}{
		{
			[][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			},
			16,
		},
	}

	for _, c := range cases {
		if got := islandPerimeter(c.in); got != c.want {
			t.Errorf(
				"\n input: %v \n got: %d \n want: %d",
				c.in,
				got,
				c.want,
			)
		}
	}
}
