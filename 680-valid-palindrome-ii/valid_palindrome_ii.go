package solution

//
// solution 1
//

func validPalindrome(s string) bool {
	var rec func(string, bool) bool
	rec = func(s string, deleted bool) bool {
		if len(s) <= 1 {
			return true
		}
		if s[0] == s[len(s)-1] {
			return rec(s[1:len(s)-1], deleted)
		}
		if deleted {
			return false
		}
		return rec(s[1:len(s)], true) || rec(s[0:len(s)-1], true)

	}
	return rec(s, false)
}

//
// solution 2
//

func validPalindrome2(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return isValid(s[i+1:j+1]) || isValid(s[i:j])
		}
	}
	return true
}

func isValid(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
