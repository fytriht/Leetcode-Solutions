package solution

import (
	"math"
)

func findUnsortedSubarray(nums []int) int {
	start, ret := math.MaxInt32, 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] <= nums[i] {
			continue
		}
		j := i
		for j > 0 && nums[j-1] > nums[j] {
			nums[j-1], nums[j] = nums[j], nums[j-1]
			j--
		}
		start = min(start, j)
		ret = max(ret, i-start+1)
	}
	return ret
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
