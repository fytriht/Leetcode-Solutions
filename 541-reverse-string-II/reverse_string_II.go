package solution

//
// solution 1
//

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

//
// solution 2
//

func reverseStr2(s string, k int) string {
	r := []rune(s)
	for i := 0; i < len(r); i += 2 * k {
		for left, right := i, min(i+k-1, len(r)-1); left < right; left, right = left+1, right-1 {
			r[left], r[right] = r[right], r[left]
		}
	}
	return string(r)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
