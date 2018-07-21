package solution

import (
	"strings"
	"unicode"
)

func mostCommonWord(paragraph string, banned []string) string {
	bm := map[string]bool{}
	for _, s := range banned {
		bm[s] = true
	}

	words := strings.FieldsFunc(paragraph, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	max := 0
	ret := ""
	wm := map[string]int{}
	for _, w := range words {
		w = strings.ToLower(w)
		if bm[w] {
			continue
		}
		wm[w]++
		if wm[w] > max {
			ret = w
			max = wm[w]
		}
	}
	return ret
}
