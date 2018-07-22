package solution

import (
	"strings"
	"unicode"
)

func toGoatLatin(S string) string {
	ret := strings.Split(S, " ")
	for i, s := range ret {
		if !isVowel(rune(s[0])) {
			s = s[1:] + string(s[0])
		}
		ret[i] = s + "ma" + strings.Repeat("a", i+1)
	}
	return strings.Join(ret, " ")
}

func isVowel(r rune) bool {
	r = unicode.ToLower(r)
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
		return true
	}
	return false
}
