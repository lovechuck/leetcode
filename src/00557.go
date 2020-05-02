package src

import "strings"

/*
557. 反转字符串中的单词 III
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

示例 1:

输入: "Let's take LeetCode contest"
输出: "s'teL ekat edoCteeL tsetnoc"
*/
func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	res := ""
	for i, s := range arr {
		res += reverseStrings(s)
		if i != len(arr)-1 {
			res += " "
		}
	}
	return res
}
func reverseStrings(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
