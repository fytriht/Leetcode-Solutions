package solution

func reverseStr(s string, k int) string {
	if k >= len(s) {
		return reverse(s)
	}
	if 2*k >= len(s) {
		return reverse(s[:k]) + s[k:]
	}
	return reverse(s[:k]) + s[k:2*k] + reverseStr(s[2*k:], k)
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
