package src

import (
	"fmt"
	"strings"
)

/*482. 密钥格式化*/

func licenseKeyFormatting(S string, K int) string {
	k := 0
	res := ""
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] != '-' {
			res = strings.ToUpper(string(S[i])) + res
			k++
			if k == K {
				k = 0
				res = "-" + res
			}
		}
	}

	return strings.Trim(res, "-")
}

func Test_licenseKeyFormatting() {
	fmt.Printf("%s", licenseKeyFormatting("2-5g-3-J", 2))
}
