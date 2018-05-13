package solution

func findDisappearedNumbers(nums []int) []int {
	m := map[int]int{}
	for i := 1; i <= len(nums); i++ {
		m[i]++
	}
	for _, n := range nums {
		m[n]--
	}

	ret := []int{}
	for k, v := range m {
		if v > 0 {
			ret = append(ret, k)
		}
	}
	return ret
}
