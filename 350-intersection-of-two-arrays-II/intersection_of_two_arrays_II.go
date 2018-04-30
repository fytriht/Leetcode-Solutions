package solution

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
