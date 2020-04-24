package src

import (
	"strings"
	"unicode"
)

/*
520. 检测大写字母

给定一个单词，你需要判断单词的大写使用是否正确。

我们定义，在以下情况时，单词的大写用法是正确的：

全部字母都是大写，比如"USA"。
单词中所有字母都不是大写，比如"leetcode"。
如果单词不只含有一个字母，只有首字母大写， 比如 "Google"。
否则，我们定义这个单词没有正确使用大写字母。
*/
func detectCapitalUse(word string) bool {
	res := strings.ToLower(word) == word || strings.ToUpper(word) == word
	if res {
		return res
	}
	rs := []rune(word)
	if !unicode.IsUpper(rs[0]) {
		return false
	}
	left := word[1:]
	return strings.ToLower(left) == left
}
