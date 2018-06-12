package solution

func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	var ret int
	for k := range m {
		if m[k+1] != 0 {
			ret = max(ret, m[k]+m[k+1])
		}
		if m[k-1] != 0 {
			ret = max(ret, m[k]+m[k-1])
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
