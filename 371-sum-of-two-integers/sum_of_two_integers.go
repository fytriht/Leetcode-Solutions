package solution

func getSum(a, b int) int {
	var ret int
	for i := 1; i <= a; i++ {
		ret++
	}
	for i := 1; i <= b; i++ {
		ret++
	}
	return ret
}
