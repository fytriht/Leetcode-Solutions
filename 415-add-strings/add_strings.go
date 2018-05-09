package solution

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	if len(num1) > len(num2) {
		return addStrings(num2, num1)
	}
	var ret string
	var carry int
	for i, j := len(num1)-1, len(num2)-1; j >= 0; j-- {
		s := int(num2[j]-'0') + carry
		if i >= 0 {
			s += int(num1[i] - '0')
			i--
		}
		if s >= 10 {
			carry = 1
			s -= 10
		} else {
			carry = 0
		}
		ret = strconv.Itoa(s) + ret
	}
	if carry == 1 {
		return "1" + ret
	}
	return ret
}
