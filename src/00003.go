package src

import "fmt"

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。


示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。


示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func lengthOfLongestSubstring(s string) int {
	start := 0

	length := 0
	result := 0

	r := []rune(s)
	hm := make(map[rune]int)

	for i := 0; i < len(r); i++ {
		if _, ok := hm[r[i]]; ok {
			if hm[r[i]] >= start {
				if length > result {
					result = length
				}
				start = hm[r[i]] + 1
				length = i - start + 1
			} else {
				length++
			}
			hm[r[i]] = i
		} else {
			hm[r[i]] = i
			length++
		}
	}
	if length > result {
		result = length
	}
	return result
}

func Test_lengthOfLongestSubstring() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("abba"))
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
	fmt.Println(lengthOfLongestSubstring("aabaab!bb"))
}
