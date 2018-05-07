package solution

import "testing"

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"Aa", 1},
		{"abccccdd", 7},
		{"a", 1},
		{"aaa", 3},
	}

	for _, c := range cases {
		if got := longestPalindrome(c.in); got != c.want {
			t.Errorf(
				"\n input: %s \n got: %d \n want: %d",
				c.in,
				got,
				c.want,
			)
		}
	}
}
