package solution

import (
	"sort"
)

//
// solution 1
//

func intersect(nums1, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	var m = make(map[int]int, len(nums1))
	var ret []int
	for _, n := range nums1 {
		m[n]++
	}
	for _, n := range nums2 {
		if m[n] != 0 {
			ret = append(ret, n)
			m[n]--
		}
	}
	return ret
}

//
// solution 2
//

func intersect2(nums1, nums2 []int) []int {
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})
	var ret []int
	var i, j int
	for i < len(nums1) && j < len(nums2) {
		switch {
		case nums1[i] < nums2[j]:
			i++
		case nums1[i] > nums2[j]:
			j++
		default:
			ret = append(ret, nums1[i])
			i++
			j++
		}
	}
	return ret
}
