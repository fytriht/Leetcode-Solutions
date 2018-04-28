package solution

import (
	"testing"
)

func TestCanWinNim(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, true},
		{7, true},
		{8, false},
		{9, true},
		{10, true},
		{11, true},
		{12, false},
	}
	for _, c := range cases {
		if got := canWinNim(c.in); got != c.want {
			t.Errorf(
				"\n input: %d \n got: %t \n want: %t",
				c.in,
				got,
				c.want,
			)
		}
	}
}
