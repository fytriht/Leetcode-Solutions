package solution

func findErrorNums(nums []int) []int {
	ret := make([]int, 2)
	for _, n := range nums {
		if nums[abs(n)-1] < 0 {
			ret[0] = abs(n)
		} else {
			nums[abs(n)-1] *= -1
		}
	}
	for i, n := range nums {
		if n > 0 {
			ret[1] = i + 1
			break
		}
	}
	return ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
