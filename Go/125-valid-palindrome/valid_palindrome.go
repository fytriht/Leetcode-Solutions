package palindrome

import (
	"unicode"
)

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isAlnum(rune(s[i])) {
			i++
		}
		for i < j && !isAlnum(rune(s[j])) {
			j--
		}
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlnum(x rune) bool {
	return unicode.IsLetter(x) || unicode.IsDigit(x)
}
