package solution

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := map[string]int{}
	for _, dm := range cpdomains {
		sl := strings.Split(dm, " ")
		cnt, _ := strconv.Atoi(sl[0])
		subdms := strings.Split(sl[1], ".")
		for i := len(subdms); i > 0; i-- {
			m[strings.Join(subdms, ".")] += cnt
			subdms = subdms[1:]
		}
	}

	ret := []string{}
	for dm, cnt := range m {
		ret = append(ret, strconv.Itoa(cnt)+" "+dm)
	}
	return ret
}
