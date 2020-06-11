package src

/*
1018. 可被 5 整除的二进制前缀

给定由若干 0 和 1 组成的数组 A。我们定义 N_i：从 A[0] 到 A[i] 的第 i 个子数组被解释为一个二进制数（从最高有效位到最低有效位）。

返回布尔值列表 answer，只有当 N_i 可以被 5 整除时，答案 answer[i] 为 true，否则为 false。
*/
func prefixesDivBy5(A []int) []bool {
	sum := 0
	var ans []bool
	for i := 0; i < len(A); i++ {
		sum = (sum<<1 + A[i]) % 25 // 注意越界
		if sum%5 == 0 {
			ans = append(ans, true)
		} else {
			ans = append(ans, false)
		}
	}
	return ans
}
