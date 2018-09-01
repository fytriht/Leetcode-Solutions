package solution

func searchRange(nums []int, target int) []int {
	pos := -1
	for left, right := 0, len(nums)-1; left <= right; {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			pos = mid
			break
		}
	}
	if pos == -1 {
		return []int{-1, -1}
	}
	left, right := pos, pos
	for left-1 >= 0 && nums[left-1] == target {
		left--
	}
	for right+1 <= len(nums)-1 && nums[right+1] == target {
		right++
	}
	return []int{left, right}
}
