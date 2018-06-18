package solution

func judgeCircle(moves string) bool {
	m := make(map[rune]int, 4)
	for _, r := range moves {
		m[r]++
	}
	return m['R'] == m['L'] && m['U'] == m['D']
}
