package solution

func isOneBitCharacter(bits []int) bool {
	var i int
	for i < len(bits)-1 {
		i += bits[i] + 1
	}
	return i == len(bits)-1
}
