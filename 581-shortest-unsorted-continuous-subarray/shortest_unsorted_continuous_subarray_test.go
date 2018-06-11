package solution

import (
	"testing"
)

func TestFindUnsortedSubarray(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 6, 4, 8, 10, 9, 15}, 5},
		{[]int{2, 3, 4, 8, 10, 9, 15}, 2},
		{[]int{2, 6, 2, 3, 4, 5, 15}, 5},
	}
	for _, c := range cases {
		if got := findUnsortedSubarray(c.in); got != c.want {
			t.Errorf(
				"\n input: %v \n want: %d \n got: %d",
				c.in,
				c.want,
				got,
			)
		}
	}
}
