package solution

func lengthOfLongestSubstring(s string) int {
	m := map[rune]int{}
	left, ret := 0, 0
	for i, r := range s {
		if v, ok := m[r]; ok && v >= left {
			left = v + 1
		}
		ret = max(ret, i-left+1)
		m[r] = i
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
