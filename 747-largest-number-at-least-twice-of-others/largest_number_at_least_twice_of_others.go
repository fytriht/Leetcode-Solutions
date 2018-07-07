package solution

import (
	"math"
)

func dominantIndex(nums []int) int {
	var idx int

	a, b := math.MinInt32, math.MinInt32
	for i, n := range nums {
		if n > a {
			a, b = n, a
			idx = i
		} else if n > b {
			b = n
		}
	}

	if a < 2*b {
		return -1
	}
	return idx
}
