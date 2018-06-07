package solution

func checkPerfectNumber(num int) bool {
	if num <= 0 {
		return false
	}

	var sum int
	for i := 1; i*i < num; i++ {
		if num%i == 0 {
			sum += (num/i + i)
		}
		if i*i == num {
			sum -= i
		}
	}
	return sum == 2*num
}
