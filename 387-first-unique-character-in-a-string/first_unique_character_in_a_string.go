package solution

func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	for i, r := range s {
		if m[r] == 1 {
			return i
		}
	}
	return -1
}
