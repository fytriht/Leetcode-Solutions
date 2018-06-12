package solution

func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	var ret int
	for k := range m {
		for _, diff := range []int{-1, 1} {
			if m[k+diff] != 0 {
				ret = max(ret, m[k]+m[k+diff])
			}
		}
	}
	return ret
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
