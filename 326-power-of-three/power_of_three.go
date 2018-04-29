package solution

//
// solution 1
//

func isPowerOfThree(n int) bool {
	for c := 1; c <= n; c *= 3 {
		if c == n {
			return true
		}
	}
	return false
}

//
// solution 2
//

func isPowerOfThree2(n int) bool {
	if n == 1 {
		return true
	}
	if n == 0 || n%3 != 0 {
		return false
	}
	return isPowerOfThree2(n / 3)
}
