package src

import (
	"fmt"
	"strings"
)

/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false
*/
func isMatch(s string, p string) bool {
	return isMatchInner(strings.Split(s, ""), strings.Split(p, ""))
}

func isMatchInner(s []string, p []string) bool {
	// dp[i][j] 表示 s[i] 在 p[j] 已经匹配
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true               //dp[i][j] 表示 s 的前 i 个是否能被 p 的前 j 个匹配
	for i := 0; i < len(p); i++ { // here's the p's length, not s's
		if p[i] == "*" && dp[0][i-1] {
			dp[0][i+1] = true // here's y axis should be i+1
		}
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if p[j] == "." || p[j] == s[i] { //如果是任意元素 或者是对于元素匹配
				dp[i+1][j+1] = dp[i][j]
			}
			if p[j] == "*" {
				if p[j-1] != s[i] && p[j-1] != "." { //如果前一个元素不匹配 且不为任意元素
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					dp[i+1][j+1] = dp[i+1][j] || dp[i][j+1] || dp[i+1][j-1]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

func Test_isMath() {
	fmt.Println(isMatch("aa", "a*"))
}
