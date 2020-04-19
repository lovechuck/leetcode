package src

import "strings"

/*459. 重复的子字符串*/
func repeatedSubstringPattern(s string) bool {
	if s == "" {
		return true
	}

	return strings.Index((s + s)[1:len(s)*2-1], s) > -1
}
