package solution

func pivotIndex(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}

	var left int
	for i, n := range nums {
		if left == sum-left-n {
			return i
		}
		left += n
	}
	return -1
}
