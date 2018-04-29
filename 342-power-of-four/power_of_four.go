package solution

func isPowerOfFour(num int) bool {
	if num == 1 {
		return true
	}
	if num == 0 || num%4 != 0 {
		return false
	}
	return isPowerOfFour(num / 4)
}
