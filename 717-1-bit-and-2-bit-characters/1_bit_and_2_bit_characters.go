package solution

func isOneBitCharacter(bits []int) bool {
	return len(bits) == 1 || bits[len(bits)-2] == 0 || !canConstruct(bits[:len(bits)-2])
}

func canConstruct(bits []int) bool {
	if len(bits) == 0 {
		return true
	}
	if len(bits) == 1 {
		return bits[0] == 0
	}
	return bits[len(bits)-1] == 0 && canConstruct(bits[:len(bits)-1]) || bits[len(bits)-2] == 1 && canConstruct(bits[:len(bits)-2])
}
