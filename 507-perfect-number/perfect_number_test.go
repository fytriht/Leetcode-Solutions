package solution

import (
	"testing"
)

func TestCheckPerfectNumber(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{-1, false},
		{0, false},
		{1, false},
		{28, true},
		{23423334233, false},
	}

	for _, c := range cases {
		if got := checkPerfectNumber(c.in); got != c.want {
			t.Errorf(
				"\n input: %d \n got: %t \n want: %t",
				c.in,
				got,
				c.want,
			)
		}
	}
}
