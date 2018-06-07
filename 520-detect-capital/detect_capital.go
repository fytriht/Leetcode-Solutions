package solution

func detectCapitalUse(word string) bool {
	var cnt int
	for _, r := range word {
		if isUpper(r) {
			cnt++
		}
	}
	return cnt == 0 || cnt == len(word) || cnt == 1 && isUpper(rune(word[0]))
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}
