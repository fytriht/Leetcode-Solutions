package solution

import (
	"testing"
)

func TestSumRange(t *testing.T) {
	nums := Constructor([]int{-2, 0, 3, -5, 2, -1})

	cases := []struct {
		ina  int
		inb  int
		want int
	}{
		{0, 2, 1},
		{2, 5, -1},
		{0, 5, -3},
	}

	for _, c := range cases {
		if got := nums.SumRange(c.ina, c.inb); got != c.want {
			t.Errorf(
				"\n input: (%d, %d) \n got: %d \n want: %d",
				c.ina,
				c.inb,
				got,
				c.want,
			)
		}
	}
}
