package solution

import (
	"strings"
)

func reverseWords(s string) string {
	sl := strings.Split(s, " ")
	for i, str := range sl {
		sl[i] = reverse(str)
	}
	return strings.Join(sl, " ")
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
