package solution

//
// solution 1
//

func isPerfectSquare(num int) bool {
	for i := 1; ; i++ {
		if i*i < num {
			continue
		}
		return i*i == num
	}
}

//
// solution 2
//

func isPerfectSquare2(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := left + (right-left)/2
		switch {
		case mid*mid > num:
			right = mid - 1
		case mid*mid < num:
			left = mid + 1
		default:
			return true
		}
	}
	return false
}
