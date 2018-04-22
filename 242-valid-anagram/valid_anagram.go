package solution

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := map[rune]int{}
	for _, b := range s {
		m[b]++
	}
	for _, b := range t {
		m[b]--
		if m[b] < 0 {
			return false
		}
	}
	return true
}
