package solution

import (
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	numsCopy := append([]int(nil), nums...)
	sort.Slice(numsCopy, func(i, j int) bool {
		return numsCopy[i] > numsCopy[j]
	})

	m := make(map[int]string)
	for i, n := range numsCopy {
		var str string
		switch i {
		case 0:
			str = "Gold Medal"
		case 1:
			str = "Silver Medal"
		case 2:
			str = "Bronze Medal"
		default:
			str = strconv.Itoa(i + 1)
		}
		m[n] = str
	}

	ret := make([]string, len(numsCopy))
	for i, n := range nums {
		ret[i] = m[n]
	}
	return ret
}
