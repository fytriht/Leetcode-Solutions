package solutoin

import (
	"testing"
)

func TestHammingDistance(t *testing.T) {
	cases := []struct {
		ina  int
		inb  int
		want int
	}{
		{1, 4, 2},
	}

	for _, c := range cases {
		if got := hammingDistance(c.ina, c.inb); got != c.want {
			t.Errorf(
				"\n input: %d, %d \n got: %d \n want: %d",
				c.ina,
				c.inb,
				got,
				c.want,
			)
		}
	}
}
