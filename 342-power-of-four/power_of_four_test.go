package solution

import (
	"testing"
)

func TestIsPowerOfFour(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{0, false},
		{1, true},
		{2, false},
		{3, false},
		{4, true},
		{65536, true},
	}
	for _, c := range cases {
		if got := isPowerOfFour(c.in); got != c.want {
			t.Errorf(
				"\n input: %d \n got: %t \n want: %t",
				c.in,
				got,
				c.want,
			)
		}
	}
}
