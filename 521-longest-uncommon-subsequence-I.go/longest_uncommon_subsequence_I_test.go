package solution

import (
	"testing"
)

func TestFindLUSlength(t *testing.T) {
	cases := []struct {
		ina  string
		inb  string
		want int
	}{
		{"a", "a", -1},
		{"aba", "cdc", 3},
		{"aba", "dsoihdsijfisjdf", 15},
	}

	for _, c := range cases {
		if got := findLUSlength(c.ina, c.inb); got != c.want {
			t.Errorf(
				"\n input: %s, %s \n got: %d \n want: %d",
				c.ina,
				c.inb,
				got,
				c.want,
			)
		}
	}
}
