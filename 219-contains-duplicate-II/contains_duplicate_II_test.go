package solution

import (
	"testing"
)

type testcase struct {
	ina  []int
	inb  int
	want bool
}

func TestContainsDuplicate(t *testing.T) {
	cases := []testcase{
		{[]int{}, 0, false},
		{[]int{}, 1, false},
		{[]int{}, 100, false},
		{[]int{1}, 0, false},
		{[]int{1}, 1, false},
		{[]int{1}, 100, false},
		{[]int{1, 1, 3, 4}, 0, false},
		{[]int{1, 1, 3, 4}, 1, true},
		{[]int{1, 1, 3, 4}, 100, true},
		{[]int{1, 3, 4, 1}, 2, false},
		{[]int{1, 1, 3, 4}, 3, true},
		{[]int{1, 1, 3, 4}, 100, true},
	}
	for _, c := range cases {
		if got := containsNearbyDuplicate(c.ina, c.inb); got != c.want {
			t.Errorf(
				"containsNearbyDuplicate(%v, %d): expected %t but got %t",
				c.ina,
				c.inb,
				c.want,
				got,
			)
		}

	}
}
