package solution

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := [][]int{}
	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			if s := nums[i] + nums[j] + nums[k]; s > 0 {
				k--
			} else if s < 0 {
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
