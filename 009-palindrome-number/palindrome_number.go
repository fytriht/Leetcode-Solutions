package solution

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	bit := 1
	for x/bit >= 10 {
		bit = bit * 10
	}
	for x != 0 {
		if x/bit != x%10 {
			return false
		}
		x, bit = x%bit/10, bit/100
	}
	return true
}
