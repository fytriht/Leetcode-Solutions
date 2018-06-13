package solution

import (
	"math"
)

func findRestaurant(list1, list2 []string) []string {
	m := map[string]int{}
	for i, s := range list2 {
		m[s] = i
	}

	var ret []string
	idxSum := math.MaxInt32
	for i, s := range list1 {
		if _, ok := m[s]; !ok {
			continue
		}
		if sum := i + m[s]; sum < idxSum {
			ret = []string{s}
			idxSum = sum
		} else if sum == idxSum {
			ret = append(ret, s)
		}
	}
	return ret
}
