package solution

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
