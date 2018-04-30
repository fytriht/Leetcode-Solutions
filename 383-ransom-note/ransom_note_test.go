package solution

import (
	"testing"
)

func TestCanConstruct(t *testing.T) {
	cases := []struct {
		ina  string
		inb  string
		want bool
	}{
		{"", "", true},
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}
	for _, c := range cases {
		if got := canConstruct(c.ina, c.inb); got != c.want {
			t.Errorf(
				"\n input: %s, %s \n got: %t \n want: %t",
				c.ina,
				c.inb,
				got,
				c.want,
			)
		}
	}
}
