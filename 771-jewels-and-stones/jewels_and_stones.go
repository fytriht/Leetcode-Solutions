package solution

func numJewelsInStones(J string, S string) int {
	m := make(map[rune]bool, len(J))
	for _, r := range J {
		m[r] = true
	}
	var cnt int
	for _, r := range S {
		if m[r] {
			cnt++
		}
	}
	return cnt
}
