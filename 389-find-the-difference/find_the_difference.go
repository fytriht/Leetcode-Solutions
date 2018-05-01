package solution

func findTheDifference(s string, t string) byte {
	var ret rune
	for _, r := range t {
		ret += r
	}
	for _, r := range s {
		ret -= r
	}
	return byte(ret)
}
