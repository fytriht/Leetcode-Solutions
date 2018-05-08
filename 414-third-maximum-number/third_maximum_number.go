package solution

import (
	"math"
)

func thirdMax(nums []int) int {
	i, j, k := math.MinInt64, math.MinInt64, math.MinInt64
	for _, n := range nums {
		switch {
		case n > i:
			k, j, i = j, i, n
		case n == i:
		case n > j:
			k, j = j, n
		case n == j:
		case n >= k:
			k = n
		}
	}
	if k == math.MinInt64 {
		return i
	}
	return k
}
