package solution

func isPerfectSquare(num int) bool {
	i := 1
	for {
		switch {
		case i*i < num:
			i++
		case i*i > num:
			return false
		default:
			return true
		}
	}
}
