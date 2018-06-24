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

func repeatedStringMatch2(A string, B string) int {
	for s, i := A, 1; i <= len(B)/len(A)+2; i++ {
		if strings.Contains(s, B) {
			return i
		}
		s += A
	}
	return -1
}
