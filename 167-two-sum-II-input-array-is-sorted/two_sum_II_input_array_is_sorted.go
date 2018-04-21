package twosum

func twoSum(numbers []int, target int) []int {
	start, end := 0, len(numbers)-1
	for start < end {
		if sum := numbers[start] + numbers[end]; sum < target {
			start++
		} else if sum > target {
			end--
		} else {
			break
		}
	}
	return []int{start + 1, end + 1}
}
