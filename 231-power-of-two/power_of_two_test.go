package solution

import (
	"testing"
)

type testcase struct {
	in   int
	want bool
}

func TestIsPowerOfTwo(t *testing.T) {
	cases := []testcase{
		{0, false},
		{1, true},
		{2, true},
		{3, false},
		{4, true},
		{5, false},
		{6, false},
		{7, false},
		{8, true},
		{1073741824, true},
		{1073741825, false},
	}
	for _, c := range cases {
		if got := isPowerOfTwo(c.in); got != c.want {
			t.Errorf("isPowerOfTwo(%d)  expected: %t, but got: %t", c.in, c.want, got)
		}
	}
}
