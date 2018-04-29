package solution

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"a", "a"},
		{"abc", "cba"},
		{"abcdefghijk", "kjihgfedcba"},
	}
	for _, c := range cases {
		if got := reverseString(c.in); got != c.want {
			t.Errorf(
				"\n input: %s \n got: %s \n want: %s",
				c.in,
				got,
				c.want,
			)
		}
	}
}
