package solution

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	ret := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k, target := i+1, len(nums)-1, -nums[i]; j < k; {
			if nums[j]+nums[k] > target {
				k--
			} else if nums[j]+nums[k] < target {
				j++
			} else {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; {
					j++
				}
				for k--; j < k && nums[k] == nums[k+1]; {
					k--
				}
			}
		}
	}
	return ret
}
