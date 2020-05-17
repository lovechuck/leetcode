package src

/*
840. 矩阵中的幻方

3 x 3 的幻方是一个填充有从 1 到 9 的不同数字的 3 x 3 矩阵，其中每行，每列以及两条对角线上的各数之和都相等。

给定一个由整数组成的 grid，其中有多少个 3 × 3 的 “幻方” 子矩阵？（每个子矩阵都是连续的）。
*/
func numMagicSquaresInside(grid [][]int) int {
	R := len(grid)
	C := len(grid[0])
	ans := 0
	for r := 0; r < R-2; r++ {
		for c := 0; c < C-2; c++ {
			if grid[r+1][c+1] != 5 {
				continue // optional skip
			}
			values := []int{grid[r][c], grid[r][c+1], grid[r][c+2],
				grid[r+1][c], grid[r+1][c+1], grid[r+1][c+2],
				grid[r+2][c], grid[r+2][c+1], grid[r+2][c+2]}
			if magic(values) {
				ans++
			}
		}
	}

	return ans
}

func magic(vals []int) bool {
	count := make([]int, 16)
	for _, val := range vals {
		count[val]++
	}
	for v := 1; v <= 9; v++ {
		if count[v] != 1 {
			return false
		}
	}

	return vals[0]+vals[1]+vals[2] == 15 &&
		vals[3]+vals[4]+vals[5] == 15 &&
		vals[6]+vals[7]+vals[8] == 15 &&
		vals[0]+vals[3]+vals[6] == 15 &&
		vals[1]+vals[4]+vals[7] == 15 &&
		vals[2]+vals[5]+vals[8] == 15 &&
		vals[0]+vals[4]+vals[8] == 15 &&
		vals[2]+vals[4]+vals[6] == 15
}
