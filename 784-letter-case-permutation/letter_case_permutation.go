package solution

import (
	"unicode"
)

//
// solution 1
//

func letterCasePermutation(S string) []string {
	ret := []string{""}
	for _, b := range S {
		if b > '9' {
			next := []string{}
			for _, s := range ret {
				next = append(next, s+string(b), s+string(b^32))
			}
			ret = next
		} else {
			for i := range ret {
				ret[i] += string(b)
			}
		}
	}
	return ret
}

//
// solution 2
//

func letterCasePermutation2(S string) []string {
	var letterCnt uint32
	for _, b := range S {
		if b > '9' {
			letterCnt++
		}
	}

	var ret []string
	for i := 0; i < 1<<letterCnt; i++ {
		var j uint32
		var s string
		for _, b := range S {
			if b > '9' {
				if (i>>j)&1 == 1 {
					s += string(unicode.ToLower(b))
				} else {
					s += string(unicode.ToUpper(b))
				}
				j++
			} else {
				s += string(b)
			}
		}
		ret = append(ret, s)
	}
	return ret
}
