package src

import "unicode"

/*
784. 字母大小写全排列

示例:
输入: S = "a1b2"
输出: ["a1b2", "a1B2", "A1b2", "A1B2"]

s[k+1]= s[k]+A & a[k]+a
*/

func letterCasePermutation(S string) []string {
	var res []string
	for _, i := range S {
		if unicode.IsLetter(i) {
			var temp []string
			if res != nil {
				for j := 0; j < len(res); j++ {
					temp = append(temp, res[j]+string(unicode.ToLower(i)))
					temp = append(temp, res[j]+string(unicode.ToUpper(i)))
				}
			} else {
				temp = append(temp, string(unicode.ToLower(i)))
				temp = append(temp, string(unicode.ToUpper(i)))
			}
			res = temp
		} else {
			if res != nil {
				for j := 0; j < len(res); j++ {
					res[j] += string(i)
				}
			} else {
				res = append(res, string(i))
			}
		}
	}
	return res
}
