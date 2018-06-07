package solution

func detectCapitalUse(word string) bool {
	return isAllUpper(word) || isAllLower(word) || isOnlyFirstLetterUpper(word)
}

func isOnlyFirstLetterUpper(s string) bool {
	return isUpper(rune(s[0])) && isAllLower(s[1:])
}

func isAllUpper(s string) bool {
	for _, r := range s {
		if isLower(r) {
			return false
		}
	}
	return true
}

func isAllLower(s string) bool {
	for _, r := range s {
		if isUpper(r) {
			return false
		}
	}
	return true
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLower(r rune) bool {
	return !isUpper(r)
}
