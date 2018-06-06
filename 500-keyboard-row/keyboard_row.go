package solution

import (
	"strings"
)

// TODO: test

var r1 = map[byte]bool{
	'q': true,
	'w': true,
	'e': true,
	'r': true,
	't': true,
	'y': true,
	'u': true,
	'i': true,
	'o': true,
	'p': true,
}

var r2 = map[byte]bool{
	'a': true,
	's': true,
	'd': true,
	'f': true,
	'g': true,
	'h': true,
	'j': true,
	'k': true,
	'l': true,
}

var r3 = map[byte]bool{
	'z': true,
	'x': true,
	'c': true,
	'v': true,
	'b': true,
	'n': true,
	'm': true,
}

func findWords(words []string) []string {
	var ret []string
	for _, w := range words {
		lower := strings.ToLower(w)

		var r map[byte]bool
		switch {
		case r1[lower[0]]:
			r = r1
		case r2[lower[0]]:
			r = r2
		case r3[lower[0]]:
			r = r3
		}

		isSameRow := true
		for _, b := range lower {
			if !r[byte(b)] {
				isSameRow = false
				break
			}
		}
		if isSameRow {
			ret = append(ret, w)
		}
	}

	return ret
}
