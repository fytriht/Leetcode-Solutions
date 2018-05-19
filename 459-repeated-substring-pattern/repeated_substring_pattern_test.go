package solution

import (
	"testing"
)

func TestRepeatedSubstringPattern(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"abab", true},
		{"aba", false},
		{"abcabcabcabc", true},
		{"bb", true},
	}

	for _, c := range cases {
		if got := repeatedSubstringPattern(c.in); got != c.want {
			t.Errorf(
				"\n input: %s \n got: %t \n want: %t",
				c.in,
				got,
				c.want,
			)
		}
	}
}
