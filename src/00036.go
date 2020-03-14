package src

/*有效的数独*/
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if board[i][j] != '.' {
				if !isValidSudokuHelp(board, i, j, board[i][j]) {
					return false
				}
			}
		}
	}
	return true
}

func isValidSudokuHelp(board [][]byte, i int, j int, c byte) bool {
	for row := 0; row < 9; row++ { //行是否合法
		if board[row][j] == c && row != i {
			return false
		}
	}
	for col := 0; col < 9; col++ { //列是否合法
		if board[i][col] == c && col != j {
			return false
		}
	}
	for row := i / 3 * 3; row < i/3*3+3; row++ { //小的block是否合法
		for col := j / 3 * 3; col < j/3*3+3; col++ {
			if board[row][col] == c && row != i && col != j {
				return false
			}
		}
	}
	return true
}
