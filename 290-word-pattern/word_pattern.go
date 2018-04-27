package solution

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	if len(strs) != len(pattern) {
		return false
	}

	m1, m2 := map[byte]string{}, map[string]byte{}
	for i, w := range strs {
		if s, ok := m1[pattern[i]]; ok {
			if w != s {
				return false
			}
		} else {
			if _, ok := m2[w]; ok {
				return false
			}
			m2[w], m1[pattern[i]] = pattern[i], w
		}
	}
	return true
}
