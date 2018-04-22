package solution

import (
	"sort"
)

//
// solution 1
//

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int, len(s))
	for _, b := range s {
		m[b]++
	}
	for _, b := range t {
		m[b]--
		if m[b] < 0 {
			return false
		}
	}
	return true
}

//
// solution 2
//

func isAnagram2(s, t string) bool {
	return SortString(s) == SortString(t)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
