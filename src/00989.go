package src

/*
989. 数组形式的整数加法

对于非负整数 X 而言，X 的数组形式是每位数字按从左到右的顺序形成的数组。例如，如果 X = 1231，那么其数组形式为 [1,2,3,1]。

给定非负整数 X 的数组形式 A，返回整数 X+K 的数组形式。
*/

func addToArrayForm(A []int, K int) []int {
	i := len(A) - 1
	cur := K
	var ans []int
	for i >= 0 || cur > 0 {
		if i >= 0 {
			cur += A[i]
		}
		ans = append(ans, cur%10)
		cur /= 10
		i--
	}
	var res []int
	for i := len(ans) - 1; i >= 0; i-- {
		res = append(res, ans[i])
	}
	return res
}
