package solution

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < nums[right] && target > nums[mid] && target <= nums[right] ||
			nums[mid] >= nums[right] && (target < nums[left] || target >= nums[mid]) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
