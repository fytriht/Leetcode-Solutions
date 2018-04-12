package rotate

func rotate(nums []int, k int) {
	for k > 0 {
		tmp := nums[len(nums)-1]
		for i := len(nums) - 2; i >= 0; i-- {
			nums[i+1] = nums[i]
		}
		nums[0] = tmp
		k--
	}
}

func rotate2(nums []int, k int) {
	curr, base := 0, 0
	for i := 0; i < len(nums); i++ {
		curr += k
		for curr > len(nums)-1 {
			curr -= len(nums)
		}
		if curr == base {
			base++
			if base > len(nums)-1 {
				base -= len(nums)
			}
			curr = base
		}
		nums[curr], nums[base] = nums[base], nums[curr]
	}
}
