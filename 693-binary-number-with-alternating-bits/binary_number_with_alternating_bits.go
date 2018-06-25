package solution

func hasAlternatingBits(n int) bool {
	for b, n := n%2, n/2; n > 0; n, b = n/2, n%2 {
		if n%2 == b {
			return false
		}
	}
	return true
}
