package solution

import (
	"sort"
	"strings"
)

//
// solution 1
//

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

//
// solution 1
//

func longestWord2(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		return strings.Compare(words[i], words[j]) < 0
	})
	var ret string
	m := make(map[string]bool, len(words))
	for _, w := range words {
		if len(w) == 1 || m[w[:len(w)-1]] {
			if len(ret) < len(w) {
				ret = w
			}
			m[w] = true
		}
	}
	return ret
}
