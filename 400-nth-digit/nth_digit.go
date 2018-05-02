package solution

import (
	"strconv"
)

func findNthDigit(n int) int {
	start, length, cnt := 1, 1, 9
	for n > length*cnt {
		n -= length * cnt
		length++
		cnt *= 10
		start *= 10
	}
	return int(strconv.Itoa(start + (n-1)/length)[(n-1)%length]) - '0'
}
