package solution

import (
	"sort"
)

func findPairs(nums []int, k int) int {
	var ret int

	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if k == 0 && nums[i+1]-nums[i] == 0 {
			ret++
		}
		for nums[i+1]-nums[i] == 0 {
			i++
			if i+1 == len(nums) {
				return ret
			}
		}

		if diff := nums[i+1] - nums[i]; diff >= k {
			if diff == k {
				ret++
			}
			continue
		}

		for j := i + 2; j < len(nums); j++ {
			if diff := nums[j] - nums[i]; diff >= k {
				if diff == k {
					ret++
				}
				break
			}
		}
	}

	return ret
}
