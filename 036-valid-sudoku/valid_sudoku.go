package solution

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if hasRepetition(board, i, i+1, 0, 9) || hasRepetition(board, 0, 9, i, i+1) {
			return false
		}
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if hasRepetition(board, i, i+3, j, j+3) {
				return false
			}
		}
	}
	return true
}

func hasRepetition(board [][]byte, ir, jr, ic, jc int) bool {
	m := make(map[byte]bool)
	for i := ir; i < jr; i++ {
		for j := ic; j < jc; j++ {
			if b := board[i][j]; b != '.' {
				if m[b] {
					return true
				}
				m[b] = true
			}
		}
	}
	return false
}
