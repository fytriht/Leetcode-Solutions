package solution

//
// solution 1
//

func validPalindrome(s string) bool {
	var rec func(string, int, int, bool) bool
	rec = func(s string, start int, end int, deleted bool) bool {
		if start >= end {
			return true
		}
		if s[start] == s[end] {
			return rec(s, start+1, end-1, deleted)
		}
		if deleted {
			return false
		}
		return rec(s, start, end-1, true) || rec(s, start+1, end, true)

	}
	return rec(s, 0, len(s)-1, false)
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
