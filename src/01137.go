package src

/*
1137. 第 N 个泰波那契数

泰波那契序列 Tn 定义如下：

T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2

给你整数 n，请返回第 n 个泰波那契数 Tn 的值。
*/
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	ans := make([]int, n+1)
	ans[0] = 0
	ans[1] = 1
	ans[2] = 1
	for i := 3; i <= n; i++ {
		ans[i] = ans[i-3] + ans[i-2] + ans[i-1]
	}
	return ans[n]
}
