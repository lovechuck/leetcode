package src

/*58. 最后一个单词的长度*/

func lengthOfLastWord(s string) int {
	sum := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			sum++
		} else if sum > 0 {
			break
		}
	}
	return sum
}
