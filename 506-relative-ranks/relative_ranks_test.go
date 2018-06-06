package solution

import (
	"testing"
)

func TestFindRelativeRanks(t *testing.T) {
	cases := []struct {
		in   []int
		want []string
	}{
		{[]int{5, 4, 3, 2, 1}, []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}},
		{[]int{2, 3, 1, 4, 5}, []string{"4", "Bronze Medal", "5", "Silver Medal", "Gold Medal"}},
	}

	for _, c := range cases {
		if got := findRelativeRanks(c.in); !isShallowEq(got, c.want) {
			t.Errorf(
				"\n input: %v \n got: %v \n want: %v",
				c.in,
				got,
				c.want,
			)
		}
	}
}

func isShallowEq(a, b []string) bool {
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
