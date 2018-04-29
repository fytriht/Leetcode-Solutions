package solution

import (
	"testing"
)

func TestReverseVowels(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"hello", "holle"},
		{"leetcode", "leotcede"},
	}
	for _, c := range cases {
		if got := reverseVowels(c.in); got != c.want {
			t.Errorf(
				"\n input: %s \n got: %s \n want: %s",
				c.in,
				got,
				c.want,
			)
		}
	}
}
