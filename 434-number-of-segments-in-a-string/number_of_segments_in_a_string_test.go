package solution

import (
	"testing"
)

func TestCountSegments(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"    a a", 2},
		{"a a    ", 2},
		{"     a a    ", 2},
		{"Hello, my name is John", 5},
	}
	for _, c := range cases {
		if got := countSegments(c.in); got != c.want {
			t.Errorf(
				"\n input: %s \n got: %d \n want: %d",
				c.in,
				got,
				c.want,
			)
		}
	}
}
