package solution

import (
	"testing"
)

func TestToHex(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{0, "0"},
		{1, "1"},
		{16, "10"},
		{-1, "ffffffff"},
		{-2, "fffffffe"},
		{26, "1a"},
		{-10, "fffffff6"},
	}
	for _, c := range cases {
		if got := toHex(c.in); got != c.want {
			t.Errorf(
				"\n input: %d \n got: %s \n want: %s",
				c.in,
				got,
				c.want,
			)
		}
	}
}
