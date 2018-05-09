package solution

//
// solution 1
//

func countSegments(s string) int {
	var ret int
	var match bool
	for _, c := range s {
		if c != ' ' {
			if !match {
				ret++
			}
			match = true
		} else {
			match = false
		}
	}
	return ret
}

//
// solution 2
//

func countSegments2(s string) int {
	var ret int
	for i := 1; i <= len(s); i++ {
		if (i == len(s) || s[i] == ' ') && s[i-1] != ' ' {
			ret++
		}
	}
	return ret
}
