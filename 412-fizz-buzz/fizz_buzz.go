package solution

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	ret := make([]string, n+1)
	for ; n > 0; n-- {
		c := ""
		if n%3 == 0 {
			c += "Fizz"
		}
		if n%5 == 0 {
			c += "Buzz"
		}
		if c == "" {
			c = strconv.Itoa(n)
		}
		ret[n] = c
	}
	return ret[1:]
}
