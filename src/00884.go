package src

import "strings"

/*
884. 两句话中的不常见单词
*/
func uncommonFromSentences(A string, B string) []string {
	hash := make(map[string]int)
	for _, s := range strings.Split(A, " ") {
		if _, ok := hash[s]; ok {
			hash[s] = hash[s] + 1
		} else {
			hash[s] = 1
		}
	}
	for _, s := range strings.Split(B, " ") {
		if _, ok := hash[s]; ok {
			hash[s] = hash[s] + 1
		} else {
			hash[s] = 1
		}
	}
	var res []string
	for s, i := range hash {
		if i == 1 {
			res = append(res, s)
		}
	}
	return res
}
