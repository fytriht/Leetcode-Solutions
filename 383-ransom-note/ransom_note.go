package solution

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, b := range magazine {
		m[b]++
	}
	for _, b := range ransomNote {
		m[b]--
		if m[b] < 0 {
			return false
		}
	}
	return true
}
