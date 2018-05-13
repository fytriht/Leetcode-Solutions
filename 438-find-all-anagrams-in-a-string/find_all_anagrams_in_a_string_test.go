package solution

import (
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	cases := []struct {
		ina  string
		inb  string
		want []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
	}

	for _, c := range cases {
		if got := findAnagrams(c.ina, c.inb); !isShallowEq(got, c.want) {
			t.Errorf(
				"\n input: %s, %s \n got: %v \n want: %v",
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
