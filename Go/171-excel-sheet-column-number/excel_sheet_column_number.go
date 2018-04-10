package title

import (
	"math"
)

func titleToNumber(s string) int {
	ret := 0
	for i := len(s) - 1; i >= 0; i-- {
		base := math.Pow(26, float64(len(s)-i-1))
		ret += int(s[i]-'A'+1) * int(base)
	}
	return ret
}
