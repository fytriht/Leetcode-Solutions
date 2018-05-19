package solution

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var i, j, ret int
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
			ret++
		}
		j++
	}
	return ret
}
