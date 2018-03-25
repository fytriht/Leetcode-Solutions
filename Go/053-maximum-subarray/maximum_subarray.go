package maxSubArray

import "math"

func maxSubArray(nums []int) int {
	m, tmp := math.MinInt32, 0
	for _, n := range nums {
		tmp += n
		m = max(tmp, m)
		if tmp < 0 {
			tmp = 0
		}
	}
	return m
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
