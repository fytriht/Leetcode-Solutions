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
	m := map[rune]int{}
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
	sort.Sort(sortRune(r))
	return string(r)
}

type sortRune []rune

func (s sortRune) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRune) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRune) Len() int {
	return len(s)
}
