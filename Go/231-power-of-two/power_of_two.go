package power

func isPowerOfTwo(n int) bool {
	for i := 1; i <= n; i *= 2 {
		if i == n {
			return true
		}
	}
	return false
}

func isPowerOfTwo2(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}
