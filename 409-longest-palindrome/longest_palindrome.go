package solution

func longestPalindrome(s string) int {
	m := map[rune]int{}
	for _, r := range s {
		m[r]++
	}

	var odd int
	for _, v := range m {
		odd += v & 1
	}

	if odd > 0 {
		odd -= 1
	}
	return len(s) - odd
}
