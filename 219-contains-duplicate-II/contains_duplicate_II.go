package solution

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i, n := range nums {
		if idx, ok := m[n]; ok && i-idx <= k {
			return true
		}
		m[n] = i
	}
	return false
}
