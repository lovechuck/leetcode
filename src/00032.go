package src

import "fmt"

/*
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
*/

func longestValidParentheses(s string) int {
	maxans := 0
	st := []rune(s)
	dp := make([]int, len(st))
	for i := 1; i < len(st); i++ {
		if st[i] == ')' {
			if st[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && st[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			if maxans < dp[i] {
				maxans = dp[i]
			}
		}
	}
	return maxans
}

func Test_longestValidParentheses() {
	fmt.Println(longestValidParentheses(")()())"))
}
