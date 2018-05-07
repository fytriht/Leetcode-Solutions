package solution

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		in   int
		want []string
	}{
		{0, []string{}},
		{1, []string{"1"}},
		{15, []string{
			"1",
			"2",
			"Fizz",
			"4",
			"Buzz",
			"Fizz",
			"7",
			"8",
			"Fizz",
			"Buzz",
			"11",
			"Fizz",
			"13",
			"14",
			"FizzBuzz",
		},
		},
	}

	for _, c := range cases {
		if got := fizzBuzz(c.in); !isShallowEq(got, c.want) {
			t.Errorf(
				"\n input: %d \n got: %v \n want: %v",
				c.in,
				got,
				c.want,
			)
		}
	}
}

func isShallowEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, n := range a {
		if n != b[i] {
			return false
		}
	}
	return true
}
