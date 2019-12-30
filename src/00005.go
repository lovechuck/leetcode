package src

import (
	"fmt"
	"strings"
)

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

*/

func longestPalindrome(s string) string {
	sa := []rune(s)
	if len(sa) <= 1 {
		return s
	}
	var arr []string
	for i := 0; i < len(sa); i++ {
		arr = append(arr, "#")
		arr = append(arr, string(sa[i]))
	}
	arr = append(arr, "#")
	length := 1
	start := 1
	end := 1
	for i := 0; i < len(arr); i = i + 1 {
		//中间向外扩散
		temp := 1
		tStart := 0
		tEnd := 0
		for j := i - 1; j >= 0 && (2*i-j) < len(arr); j-- {
			if arr[j] == arr[2*i-j] {
				temp += 2
				tStart = j
				tEnd = 2*i - j
			} else {
				break
			}
		}
		if temp > length {
			length = temp
			start = tStart
			end = tEnd
		}
	}
	var builder strings.Builder
	for i := start; i <= end; i++ {
		if arr[i] != "#" {
			builder.WriteString(arr[i])
		}
	}
	return builder.String()
}

func Test_longestPalindrome() {
	fmt.Println(longestPalindrome("aa"))
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}
