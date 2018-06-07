package solution

import (
	"testing"
)

func TestDetectCapitalUse(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"USA", true},
		{"leetcode", true},
		{"Google", true},
		{"FlaG", false},
	}

	for _, c := range cases {
		if got := detectCapitalUse(c.in); got != c.want {
			t.Errorf(
				"\n input: %s \n got: %t \n want: %t",
				c.in,
				got,
				c.want,
			)
		}
	}
}
