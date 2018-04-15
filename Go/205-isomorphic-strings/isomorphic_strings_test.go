package isomorphic

import (
	"testing"
)

type testcase struct {
	ina  string
	inb  string
	want bool
}

func TestIsIsomorphic(t *testing.T) {
	cases := []testcase{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"ab", "cc", false},
	}
	for _, c := range cases {
		if got := isIsomorphic2(c.ina, c.inb); got != c.want {
			t.Errorf(
				"isIsomorphic(\"%s\", \"%s\"): expected %t but got %t",
				c.ina,
				c.inb,
				c.want,
				got,
			)
		}
	}
}
