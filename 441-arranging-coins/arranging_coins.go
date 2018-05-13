package solution

func arrangeCoins(n int) int {
	var i int
	for sum := 0; sum <= n; i++ {
		sum += i
	}
	return i - 2
}
