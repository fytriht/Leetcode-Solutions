package solution

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
