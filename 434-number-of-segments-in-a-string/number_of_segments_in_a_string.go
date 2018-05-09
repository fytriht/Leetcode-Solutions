package solution

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
