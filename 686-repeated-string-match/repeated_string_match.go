package solution

import (
	"strings"
)

func repeatedStringMatch(A string, B string) int {
	var times int
	var s string
	for len(s) < len(B) {
		times++
		s += A
	}
	if strings.Contains(s, B) {
		return times
	}
	if strings.Contains(s+A, B) {
		return times + 1
	}
	return -1
}
