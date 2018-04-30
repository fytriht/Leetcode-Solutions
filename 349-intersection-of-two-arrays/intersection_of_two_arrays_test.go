package solution

import (
	"testing"
)

func TestIntersection(t *testing.T) {
	cases := []struct {
		ina  []int
		inb  []int
		want []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{1}, []int{}},
		{[]int{1}, []int{1}, []int{1}},
		{[]int{1}, []int{1, 1, 2, 2, 3, 3}, []int{1}},
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{}},
	}
	for _, c := range cases {
		if got := intersection(c.ina, c.inb); !isShallowEq(got, c.want) {
			t.Errorf(
				"\n input: %v, %v \n got: %v \n want: %v",
				c.ina,
				c.inb,
				got,
				c.want,
			)
		}
	}
}

func isShallowEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, n := range a {
		if n != b[i] {
			return false
		}
	}
	return true
}
