package solution

import (
	"testing"
)

func TestIsPowerOfThree(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{0, false},
		{1, true},
		{2, false},
		{3, true},
		{4, false},
		{177147, true},
	}
	for _, c := range cases {
		if got := isPowerOfThree2(c.in); got != c.want {
			t.Errorf(
				"\n input: %d \n got: %t \n want: %t",
				c.in,
				got,
				c.want,
			)
		}
	}
}
