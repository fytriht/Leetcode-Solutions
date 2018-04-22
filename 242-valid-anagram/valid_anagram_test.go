package solution

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		ina, inb string
		want     bool
	}{
		{"", "", true},
		{"aa", "aa", true},
		{"anagram", "nagaram", true},
		{"anagram", "nagaram!", false},
		{"rat", "car", false},
	}

	for _, c := range cases {
		if got := isAnagram(c.ina, c.inb); got != c.want {
			t.Errorf(
				"isAnagram(\"%s\", \"%s\"): expected %t but got %t",
				c.ina,
				c.inb,
				c.want,
				got,
			)
		}
	}
}
