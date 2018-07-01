package solution

import (
	"strings"
)

func longestWord(words []string) string {
	q := []string{}
	m := make(map[int][]string, len(words))
	for _, w := range words {
		if _, ok := m[len(w)]; ok {
			m[len(w)] = append(m[len(w)], w)
		} else {
			m[len(w)] = []string{w}
		}
		if len(w) == 1 {
			q = append(q, w)
		}
	}

	var ret string
	for len(q) > 0 {
		word := q[0]
		q = q[1:]
		if len(word) == len(ret) {
			ret = minStr(word, ret)
		} else {
			ret = word
		}
		for _, w := range m[len(word)+1] {
			if w[:len(w)-1] == word {
				q = append(q, w)
			}
		}
	}
	return ret
}

func minStr(x, y string) string {
	if strings.Compare(x, y) < 0 {
		return x
	}
	return y
}
