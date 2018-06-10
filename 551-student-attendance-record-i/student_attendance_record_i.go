package solution

func checkRecord(s string) bool {
	var as, ls int
	for _, r := range s {
		if r == 'A' {
			as++
			if as > 1 {
				return false
			}
		}
		if r == 'L' {
			ls++
			if ls > 2 {
				return false
			}
		} else {
			ls = 0
		}
	}
	return true
}
