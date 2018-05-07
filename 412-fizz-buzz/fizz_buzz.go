package solution

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	ret := []string{}
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			ret = append(ret, "FizzBuzz")
		case i%3 == 0:
			ret = append(ret, "Fizz")
		case i%5 == 0:
			ret = append(ret, "Buzz")
		default:
			ret = append(ret, strconv.Itoa(i))
		}
	}
	return ret
}
