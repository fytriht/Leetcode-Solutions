package solution

import (
	"testing"
)

func TestGetSum(t *testing.T) {
	cases := []struct {
		ina  int
		inb  int
		want int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{3, 4, 7},
	}
	for _, c := range cases {
		if got := getSum(c.ina, c.inb); got != c.want {
			t.Errorf(
				"\n input: %d %d \n got: %d \n want: %d",
				c.ina,
				c.inb,
				got,
				c.want,
			)
		}
	}
}
