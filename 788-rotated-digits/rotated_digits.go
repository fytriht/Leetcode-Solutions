package solution

import (
	"strconv"
)

//
// solution 1
//

// 0: 3, 4, 7
// 1: 0, 1, 8
// 2: 2, 5, 6, 9
func rotatedDigits(N int) int {
	var ret int

	dp := make([]int, N+1)
	for i := 0; i <= N; i++ {
		if i < 10 {
			if i == 0 || i == 1 || i == 8 {
				dp[i] = 1
			} else if i == 2 || i == 5 || i == 6 || i == 9 {
				dp[i] = 2
				ret++
			}
		} else {
			a, b := dp[i/10], dp[i%10]
			if a == 1 && b == 1 {
				dp[i] = 1
			} else if a >= 1 && b >= 1 {
				dp[i] = 2
				ret++
			}
		}
	}
	return ret
}

//
// solution 2
//

func rotatedDigits2(N int) int {
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
