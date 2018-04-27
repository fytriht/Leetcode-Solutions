package solution

import (
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 0}, []int{1, 2, 3, 0}},
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	}

	for _, c := range cases {
		in := append([]int{}, c.in...)
		moveZeroes2(c.in)
		if !isShallowEq(c.in, c.want) {
			t.Errorf(
				"\n input: %v\n expected: %v\n got: %v",
				in,
				c.want,
				c.in,
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
