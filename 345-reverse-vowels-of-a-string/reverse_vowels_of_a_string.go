package solution

func reverseVowels(s string) string {
	ret := []byte(s)
	i, j := 0, len(s)-1
	for {
		for i < j && !isVowels(s[i]) {
			i++
		}
		for i < j && !isVowels(s[j]) {
			j--
		}
		if i >= j {
			break
		}
		ret[i], ret[j] = s[j], s[i]
		i++
		j--

	}
	return string(ret)
}

func isVowels(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
