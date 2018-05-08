package solution

import (
	"testing"
)

func TestThirdMax(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{3, 2, 1}, 1},
		{[]int{2, 1}, 2},
		{[]int{2, 2, 3, 1}, 1},
		{[]int{2, 2, 2, 1}, 2},
		{[]int{2, 2, 2, 4}, 4},
		{[]int{1, -2147483648, 2}, -2147483648},
		{[]int{1, 2, -2147483648}, -2147483648},
	}
	for _, c := range cases {
		if got := thirdMax(c.in); got != c.want {
			t.Errorf(
				"\n input: %v \n got: %d \n want: %d",
				c.in,
				got,
				c.want,
			)
		}
	}
}
