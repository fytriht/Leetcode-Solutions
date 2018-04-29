package solution

//
// solution 1
//

func isPowerOfFour(num int) bool {
	if num == 1 {
		return true
	}
	if num == 0 || num%4 != 0 {
		return false
	}
	return isPowerOfFour(num / 4)
}

//
// solution 2
//

func isPowerOfFour2(num int) bool {
	for c := 1; c <= num; c *= 4 {
		if c == num {
			return true
		}
	}
	return false
}
