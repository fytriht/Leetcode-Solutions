package solution

import (
	"testing"
)

type testcase struct {
	in   []int
	want bool
}

func TestContainsDuplicate(t *testing.T) {
	cases := []testcase{
		{[]int{}, false},
		{[]int{1}, false},
		{[]int{1, 1}, true},
		{[]int{1, 1, 1}, true},
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},
	}
	for _, c := range cases {
		if got := containsDuplicate(c.in); got != c.want {
			t.Errorf(
				"containsDuplicate(%v): expected %t but got %t",
				c.in,
				c.want,
				got,
			)
		}
		if got := containsDuplicate2(c.in); got != c.want {
			t.Errorf(
				"containsDuplicate(%v): expected %t but got %t",
				c.in,
				c.want,
				got,
			)
		}

	}
}
