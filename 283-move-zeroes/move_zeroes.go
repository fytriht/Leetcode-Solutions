package solution

func moveZeroes(nums []int) {
	j := 0
	for i, n := range nums {
		if n != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
