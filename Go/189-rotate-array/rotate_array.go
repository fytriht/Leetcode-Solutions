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
		curr %= len(nums)
		if curr == base {
			base++
			base %= len(nums)
			curr = base
		}
		nums[curr], nums[base] = nums[base], nums[curr]
	}
}

func rotate3(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
