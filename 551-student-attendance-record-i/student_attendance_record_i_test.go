package solution

import (
	"testing"
)

func TestCheckRecord(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"PPAP", true},
		{"APAP", false},
		{"PPLL", true},
		{"PPLLL", false},
		{"LPLL", true},
	}

	for _, c := range cases {
		if got := checkRecord(c.in); got != c.want {
			t.Errorf(
				"\n input: %s \n got: %t \n want: %t",
				c.in,
				got,
				c.want,
			)
		}
	}
}
