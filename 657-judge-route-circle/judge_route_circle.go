package solution

//
// solution 1
//

func judgeCircle(moves string) bool {
	m := make(map[rune]int, 4)
	for _, r := range moves {
		m[r]++
	}
	return m['R'] == m['L'] && m['U'] == m['D']
}

//
// solution 2
//

func judgeCircle2(moves string) bool {
	var x, y int
	for _, r := range moves {
		switch r {
		case 'R':
			x++
		case 'L':
			x--
		case 'U':
			y++
		case 'D':
			y--
		}
	}
	return x == 0 && y == 0
}
