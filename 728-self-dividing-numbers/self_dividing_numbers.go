package solution

func selfDividingNumbers(left int, right int) []int {
	ret := []int{}
	for ; left <= right; left++ {
		if isSelfDividing(left) {
			ret = append(ret, left)
		}
	}
	return ret
}

func isSelfDividing(n int) bool {
	for m := n; n != 0; n /= 10 {
		if mod := n % 10; mod == 0 || m%mod != 0 {
			return false
		}
	}
	return true
}
