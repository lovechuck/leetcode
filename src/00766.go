package src

/*
766. 托普利茨矩阵
*/
func isToeplitzMatrix(matrix [][]int) bool {
	if matrix == nil {
		return true
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i > 0 && j > 0 && matrix[i-1][j-1] != matrix[i][j] {
				return false
			}
		}
	}
	return true
}
