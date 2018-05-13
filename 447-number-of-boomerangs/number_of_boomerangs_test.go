package solution

import (
	"testing"
)

func TestNumberOfBoomerangs(t *testing.T) {
	cases := []struct {
		in   [][]int
		want int
	}{
		{[][]int{{0, 0}, {1, 0}, {2, 0}}, 2},
	}

	for _, c := range cases {
		if got := numberOfBoomerangs(c.in); got != c.want {
			t.Errorf(
				"\n input: %v \n got: %d \n want: %d",
				c.in,
				got,
				c.want,
			)
		}
	}
}
