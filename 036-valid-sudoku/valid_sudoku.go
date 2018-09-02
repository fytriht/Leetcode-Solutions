package solution

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		rm := make(map[byte]bool)
		cm := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			rb, cb := board[i][j], board[j][i]
			if rm[rb] || cm[cb] {
				return false
			}
			rm[rb], cm[cb] = rb != '.', cb != '.'
		}
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			m := make(map[byte]bool)
			for k := 0; k < 3; k++ {
				for n := 0; n < 3; n++ {
					if b := board[i+k][j+n]; b != '.' {
						if m[b] {
							return false
						}
						m[b] = true
					}
				}

			}
		}
	}
	return true
}
