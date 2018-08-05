package solution

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ret, minDiff := 0, math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if diff := abs(sum - target); diff < minDiff {
				ret, minDiff = sum, diff
			}
			if sum > target {
				k--
			} else if sum < target {
				j++
			} else {
				return sum
			}
		}
	}
	return ret
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
