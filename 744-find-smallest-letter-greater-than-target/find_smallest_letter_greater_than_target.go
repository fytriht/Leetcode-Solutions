package solution

func nextGreatestLetter(letters []byte, target byte) byte {
	if letters[len(letters)-1] <= target {
		return letters[0]
	}
	left, right := 0, len(letters)-1
	for left < right {
		if mid := left + (right-left)/2; target >= letters[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return letters[left]
}
