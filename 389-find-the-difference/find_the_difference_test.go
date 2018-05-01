package solution

import (
	"testing"
)

func TestFindTheDifference(t *testing.T) {
	cases := []struct {
		ina  string
		inb  string
		want byte
	}{
		{"", "a", byte('a')},
		{"abcd", "ebacd", byte('e')},
	}

	for _, c := range cases {
		if got := findTheDifference(c.ina, c.inb); got != c.want {
			t.Errorf(
				"\n input: %s, %s \n got: %s \n want: %s",
				c.ina,
				c.inb,
				string(got),
				string(c.want),
			)
		}
	}
}
