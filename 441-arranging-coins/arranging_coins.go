package solution

//
// solution 1
//

func arrangeCoins(n int) int {
	var i int
	for sum := 0; sum <= n; i++ {
		sum += i
	}
	return i - 2
}

//
// solution 2
//

func arrangeCoins2(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		if 2*n < mid*(mid+1) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left - 1
}
