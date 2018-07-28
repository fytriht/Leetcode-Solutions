package solution

func lengthOfLongestSubstring(s string) int {
	ret := 0
	for i := 0; i < len(s); i++ {
		m := map[byte]bool{s[i]: true}
		j := i + 1
		for ; j < len(s) && !m[s[j]]; j++ {
			m[s[j]] = true
		}
		ret = max(ret, j-i)

	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
