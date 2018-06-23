package solution

import (
	"strconv"
)

func calPoints(ops []string) int {
	var s []int
	var ret int
	for _, op := range ops {
		var score int
		switch op {
		case "+":
			score += s[len(s)-1]
			score += s[len(s)-2]
		case "D":
			score = 2 * s[len(s)-1]
		case "C":
			ret -= s[len(s)-1]
			s = s[:len(s)-1]
			continue
		default:
			score, _ = strconv.Atoi(op)
		}
		ret += score
		s = append(s, score)
	}
	return ret
}
