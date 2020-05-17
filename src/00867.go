package src

/*867. 转置矩阵*/

func transpose(A [][]int) [][]int {
	R := len(A)
	C := len(A[0])
	ans := make([][]int, C)
	for i := 0; i < C; i++ {
		ans[i] = make([]int, R)
	}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			ans[j][i] = A[i][j]
		}
	}
	return ans
}
