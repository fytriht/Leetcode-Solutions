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

func moveZeroes2(nums []int) {
	j := 0
	for _, n := range nums {
		if n != 0 {
			nums[j] = n
			j++
		}
	}
	for j < len(nums) {
		nums[j] = 0
		j++
	}
}
