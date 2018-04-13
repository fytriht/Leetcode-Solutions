package happy

func isHappy(n int) bool {
	m := make(map[int]bool)
	for {
		n = getSumOfDigitsSuqare(n)
		if m[n] {
			return false
		}
		if n == 1 {
			return true
		}
		m[n] = true
	}
}

func getSumOfDigitsSuqare(n int) int {
	s := 0
	for n != 0 {
		s += (n % 10) * (n % 10)
		n /= 10
	}
	return s
}
