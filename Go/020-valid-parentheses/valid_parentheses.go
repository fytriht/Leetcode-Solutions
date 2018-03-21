package isValid

func isValid(s string) bool {
	var m = map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []rune{}
	for _, r := range s {
		if v, ok := m[r]; ok {
			if len(stack) == 0 || v != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}
	return len(stack) == 0
}
