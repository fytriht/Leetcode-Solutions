package isPalindrome

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	bit := 1
	for x/bit >= 10 {
		bit = bit * 10
	}
	for x != 0 {
		left := x / bit
		right := x % 10
		if left != right {
			return false
		}
		x = (x % bit) / 10
		bit = bit / 100
	}
	return true
}
