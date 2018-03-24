package searchInsert

func searchInsert(nums []int, target int) int {
	for i, n := range nums {
		if n == target || n > target {
			return i
		}
	}
	return len(nums)
}
