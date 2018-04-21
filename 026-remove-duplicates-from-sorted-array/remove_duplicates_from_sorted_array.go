package solution

func removeDuplicates(nums []int) int {
	i := -1
	for _, n := range nums {
		if i == -1 || n != nums[i] {
			i++
			nums[i] = n
		}
	}
	return i + 1
}
