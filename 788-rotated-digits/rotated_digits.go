package solution

import (
	"strconv"
)

func rotatedDigits(N int) int {
	var ret int
	for i := 1; i <= N; i++ {
		if isGoodNum(i) {
			ret++
		}
	}
	return ret
}

var m = map[byte]byte{
	'0': '0',
	'1': '1',
	'8': '8',
	'2': '5',
	'5': '2',
	'6': '9',
	'9': '6',
}

func isGoodNum(n int) bool {
	var s string
	for _, r := range strconv.Itoa(n) {
		if _, ok := m[byte(r)]; !ok {
			return false
		}
		s += string(m[byte(r)])
	}
	return s != strconv.Itoa(n)
}
