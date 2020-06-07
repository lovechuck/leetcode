package src

/*
给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）组成的列表。例如，如果一个字符在每个字符串中出现 3 次，但不是 4 次，则需要在最终答案中包含该字符 3 次。

你可以按任意顺序返回答案。
*/

func commonChars(A []string) []string {
	if len(A) == 1 {
		return A
	}

	res := getNums(A[0])
	for i := 1; i < len(A); i++ {
		temp := getNums(A[i])
		for i := 0; i < 26; i++ {
			if temp[i] < res[i] {
				res[i] = temp[i]
			}
		}
	}

	var ans []string
	for i := 0; i < 26; i++ {
		for res[i] > 0 {
			ans = append(ans, string(i+'a'))
			res[i]--
		}
	}

	return ans
}

func getNums(A string) [26]int {
	ans := [26]int{}
	for i := 0; i < len(A); i++ {
		ans[A[i]-'a']++
	}
	return ans
}
