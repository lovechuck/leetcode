package src

/*
521. 最长特殊序列 Ⅰ
题目读不懂 垃圾题
*/

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	if len(a) >= len(b) {
		return len(a)
	}

	return len(b)
}
