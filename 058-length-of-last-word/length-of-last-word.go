package lengthOfLastWord

func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if length > 0 {
				break
			}
		} else {
			length++
		}
	}
	return length
}
