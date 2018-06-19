package solution

func checkPossibility(nums []int) bool {
	var modified bool
	for i := 0; i+1 < len(nums); i++ {
		if nums[i] > nums[i+1] {
			if modified || i > 0 && i+2 < len(nums) && nums[i+1] < nums[i-1] && nums[i+2] < nums[i] {
				return false
			}
			modified = true
		}
	}
	return true
}
