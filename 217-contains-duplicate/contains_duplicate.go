package duplicate

import (
	"sort"
)

func containsDuplicate(nums []int) bool {
	m := map[int]bool{}
	for _, n := range nums {
		if m[n] {
			return true
		}
		m[n] = true
	}
	return false
}

func containsDuplicate2(nums []int) bool {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
