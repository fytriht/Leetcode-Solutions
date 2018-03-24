package searchInsert

func searchInsert(nums []int, target int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == target || target < nums[i] {
			break
		}
		i++

	}
	return i
}
