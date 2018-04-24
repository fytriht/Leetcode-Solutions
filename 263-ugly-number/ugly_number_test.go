package solution

import (
	"testing"
)

func TestIsUgly(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{0, false},
		{1, true},
		{6, true},
		{8, true},
		{14, false},
	}

	for _, c := range cases {
		if got := isUgly(c.in); got != c.want {
			t.Errorf(
				"isUgly(%d): expected %t but got %t",
				c.in,
				c.want,
				got,
			)
		}
	}
}
