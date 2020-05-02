package src

import "strings"

/*551. 学生出勤记录 I*/

func checkRecord(s string) bool {
	if strings.Index(s, "LLL") > -1 {
		return false
	}
	count := 0
	for _, i := range s {
		if i == 'A' {
			count++
		}
		if count >= 2 {
			return false
		}
	}
	return true
}
