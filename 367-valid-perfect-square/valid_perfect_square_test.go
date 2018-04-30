package solution

import (
	"testing"
)

func TestIsPerfectSquare(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{1, true},
		{2, false},
		{3, false},
		{4, true},
		{5, false},
		{14, false},
		{16, true},
		{2394232398, false},
	}
	for _, c := range cases {
		if got := isPerfectSquare(c.in); got != c.want {
			t.Errorf(
				"\n input: %d \n got: %t \n want: %t",
				c.in,
				got,
				c.want,
			)
		}
	}
}
