package solution

func isPowerOfThree(n int) bool {
	for c := 1; c <= n; c *= 3 {
		if c == n {
			return true
		}
	}
	return false
}
