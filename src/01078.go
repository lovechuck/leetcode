package src

import (
	"strings"
)

/*
1078. Bigram 分词

给出第一个词 first 和第二个词 second，考虑在某些文本 text 中可能以 "first second third" 形式出现的情况，其中 second 紧随 first 出现，third 紧随 second 出现。

对于每种这样的情况，将第三个词 "third" 添加到答案中，并返回答案。
*/

func findOcurrences(text string, first string, second string) []string {
	words := strings.Split(text, " ")
	var ans []string
	for i := 0; i < len(words); i++ {
		if i+1 < len(words) && i+2 < len(words) && words[i] == first && words[i+1] == second {
			ans = append(ans, words[i+2])
		}
	}
	return ans
}
