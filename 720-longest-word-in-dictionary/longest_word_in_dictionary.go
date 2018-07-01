package solution

import (
	"strings"
)

func longestWord(words []string) string {
	q := []string{}
	m := make(map[string]bool, len(words))
	for _, w := range words {
		m[w] = true
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
		for i := 'a'; i <= 'z'; i++ {
			if w := word + string(i); m[w] {
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
