package solution

import (
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"", -1},
		{"a", 0},
		{"abbcddeea", 3},
		{"abcde", 0},
	}
	for _, c := range cases {
		if got := firstUniqChar(c.in); got != c.want {
			t.Errorf(
				"\n input: %s \n got: %d \n want: %d",
				c.in,
				got,
				c.want,
			)
		}
	}
}
