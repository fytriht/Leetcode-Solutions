package solution

func findPairs(nums []int, k int) int {
	var ret int

	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	for key, count := range m {
		if k == 0 && count > 1 || k > 0 && m[key+k] > 0 {
			ret++
		}
	}
	return ret
}
