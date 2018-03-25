package plusOne

func plusOne(digits []int) []int {
	carry := true
	for i := len(digits) - 1; carry && i >= 0; i-- {
		digits[i] = digits[i] + 1
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			carry = false
		}
	}
	if carry {
		digits = append([]int{1}, digits...)
	}
	return digits
}
