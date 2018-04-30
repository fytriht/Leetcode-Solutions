package solution

import (
	"sort"
)

//
// solution 1
//

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersection(nums2, nums1)
	}
	m := make(map[int]bool, len(nums1))
	for _, n := range nums1 {
		m[n] = true
	}
	var ret []int
	for _, n := range nums2 {
		if m[n] {
			delete(m, n)
			ret = append(ret, n)
		}
	}
	return ret
}

//
// solution 2
//

func intersection2(nums1, nums2 []int) []int {
	for _, s := range [][]int{nums1, nums2} {
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
	}
	var ret []int
	var i, j int
	for i < len(nums1) && j < len(nums2) {
		switch {
		case nums1[i] < nums2[j]:
			i++
		case nums1[i] > nums2[j]:
			j++
		default:
			if len(ret) == 0 || ret[len(ret)-1] != nums1[i] {
				ret = append(ret, nums1[i])
			}
			i++
			j++
		}
	}
	return ret
}
