package solution

// TODO: KMP Algorithm http://en.wikipedia.org/wiki/Knuth-Morris-Pratt_algorithm

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
