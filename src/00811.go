package src

import (
	"fmt"
	"strconv"
	"strings"
)

/*
811. 子域名访问计数
*/
func subdomainVisits(cpdomains []string) []string {
	hash := make(map[string]int)
	for _, cpdomain := range cpdomains {
		arr := strings.Split(cpdomain, " ")
		num, _ := strconv.Atoi(arr[0])
		domain := arr[1]
		if _, ok := hash[domain]; ok {
			hash[domain] = hash[domain] + num
		} else {
			hash[domain] = num
		}
		for strings.Index(domain, ".") > -1 {
			domain = domain[strings.Index(domain, ".")+1:]
			if _, ok := hash[domain]; ok {
				hash[domain] = hash[domain] + num
			} else {
				hash[domain] = num
			}
		}
	}
	var res []string
	for s, i := range hash {
		res = append(res, fmt.Sprintf("%d %s", i, s))
	}
	return res
}
