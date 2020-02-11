package src

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
*/

func isValid(s string) bool {
	var stack []rune
	for _, v := range s {
		if isLeft(v) {
			stack = append(stack, v)
		} else if stack == nil || len(stack) == 0 || !isRightEqual(v, stack[len(stack)-1]) {
			return false
		} else {
			stack = stack[0 : len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func isLeft(s rune) bool {
	if s == '(' || s == '{' || s == '[' {
		return true
	}
	return false
}

func isRightEqual(r rune, l rune) bool {
	if r == ')' && l == '(' {
		return true
	} else if r == '}' && l == '{' {
		return true
	} else if r == ']' && l == '[' {
		return true
	}
	return false
}

func Test_isValid() {
	fmt.Println(isValid("[])"))
}
