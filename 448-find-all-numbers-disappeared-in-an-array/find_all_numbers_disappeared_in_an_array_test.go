package solution

import (
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{[]int{1}, []int{}},
		{[]int{1, 1}, []int{2}},
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{6, 5}},
	}

	for _, c := range cases {
		if got := findDisappearedNumbers(c.in); !isShallowEq(got, c.want) {
			t.Errorf(
				"\n input: %v \n got: %v \n want: %v",
				c.in,
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
