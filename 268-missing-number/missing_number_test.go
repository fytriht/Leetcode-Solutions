package solution

import (
	"testing"
)

func TestMissingNumber(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{0, 1}, 2},
		{[]int{3, 0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}

	for _, c := range cases {
		if got := missingNumber(c.in); got != c.want {
			t.Errorf(
				"missingNumber(%v): expected %d but got %d",
				c.in,
				c.want,
				got,
			)
		}
	}
}
