package solution

//
// solution 1
//

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

//
// solution 2
//

func searchRange2(nums []int, target int) []int {
	ret := []int{-1, -1}
	left, right := 0, len(nums)-1
	for left < right {
		if mid := left + (right-left)/2; nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if right < 0 || nums[right] != target {
		return ret
	}
	ret[0] = right
	right = len(nums)
	for left < right {
		if mid := left + (right-left)/2; nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	ret[1] = left - 1
	return ret
}
