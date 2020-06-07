package src

/*
942. 增减字符串匹配


给定只含 "I"（增大）或 "D"（减小）的字符串 S ，令 N = S.length。

返回 [0, 1, ..., N] 的任意排列 A 使得对于所有 i = 0, ..., N-1，都有：

如果 S[i] == "I"，那么 A[i] < A[i+1]
如果 S[i] == "D"，那么 A[i] > A[i+1]
*/

func diStringMatch(S string) []int {
	min := 0
	max := len(S)
	var ans []int
	for _, i := range S {
		if i == 'I' {
			ans = append(ans, min)
			min++
		} else {
			ans = append(ans, max)
			max--
		}
	}
	ans = append(ans, min)
	return ans
}
