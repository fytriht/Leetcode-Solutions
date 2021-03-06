package solution

//
// solution 1
//

func findTheDifference(s, t string) byte {
	var ret rune
	for _, r := range t {
		ret += r
	}
	for _, r := range s {
		ret -= r
	}
	return byte(ret)
}

//
// solution 2
//

func findTheDifference2(s, t string) byte {
	var ret rune
	for _, r := range s + t {
		ret ^= r
	}
	return byte(ret)
}
