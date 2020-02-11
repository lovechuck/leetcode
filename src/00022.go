package src

import "fmt"

/*
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

func generateParenthesis(n int) []string {
	var ans []string
	backtrack(&ans, "", 0, 0, n)
	return ans
}

func backtrack(ans *[]string, cur string, open int, close int, max int) {
	if len(cur) == max*2 {
		*ans = append(*ans, cur)
		return
	}

	if open < max {
		backtrack(ans, cur+"(", open+1, close, max)
	}

	if close < open {
		backtrack(ans, cur+")", open, close+1, max)
	}
}

func Test_generateParenthesis() {
	fmt.Sprint(generateParenthesis(4))
}
