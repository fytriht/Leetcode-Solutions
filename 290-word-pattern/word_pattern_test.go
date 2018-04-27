package solution

import (
	"testing"
)

func TestWordPattern(t *testing.T) {
	cases := []struct {
		ina  string
		inb  string
		want bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
	}

	for _, c := range cases {
		if got := wordPattern(c.ina, c.inb); got != c.want {
			t.Errorf(
				"\nin: %s, %s \ngot: %t \nwant: %t",
				c.ina,
				c.inb,
				got,
				c.want,
			)
		}
	}
}
