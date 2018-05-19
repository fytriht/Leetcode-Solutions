package solution

import (
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	cases := []struct {
		ina  []int
		inb  []int
		want int
	}{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
		{[]int{1, 2, 3}, []int{3, 2, 1}, 3},
		{[]int{10, 9, 8, 7}, []int{5, 6, 7, 8}, 2},
	}

	for _, c := range cases {
		if got := findContentChildren(c.ina, c.inb); got != c.want {
			t.Errorf(
				"\n input: %v, %v \n got: %d \n want: %d",
				c.ina,
				c.inb,
				got,
				c.want,
			)
		}
	}
}
