package solutoin

func hammingDistance(x, y int) int {
	var ret int
	for z := x ^ y; z != 0; z &= (z - 1) {
		ret++
	}
	return ret
}
