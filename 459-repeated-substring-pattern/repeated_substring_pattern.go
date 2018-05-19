package solution

import (
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	for i := len(s) / 2; i > 0; i-- {
		if len(s)%i == 0 && strings.Repeat(s[:i], len(s)/i) == s {
			return true
		}
	}
	return false
}
