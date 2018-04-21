package removeElement

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)
	for i < j {
		if nums[i] == val {
			j--
			nums[i] = nums[j]
		} else {
			i++
		}
	}
	return j
}
