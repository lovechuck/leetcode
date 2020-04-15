package src

import "strings"

/*434. 字符串中的单词数*/

func countSegments(s string) int {
	res := 0
	arr := strings.Split(s, " ")
	for _, item := range arr {
		item = strings.Trim(item, " ")
		if item != "" {
			res++
		}
	}
	return res
}
