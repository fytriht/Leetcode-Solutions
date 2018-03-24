package removeElement

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)
	for i < j {
		if nums[i] != val {
			i++
		} else {
			j--
			nums[i] = nums[j]
		}
	}
	return j
}
