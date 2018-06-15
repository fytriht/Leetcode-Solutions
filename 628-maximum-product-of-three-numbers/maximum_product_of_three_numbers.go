package solution

import (
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	ret := nums[len(nums)-3] * nums[len(nums)-2] * nums[len(nums)-1]
	if neg := nums[0] * nums[1] * nums[len(nums)-1]; neg > ret {
		return neg
	}
	return ret
}
