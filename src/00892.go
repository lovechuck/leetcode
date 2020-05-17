package src

/*
892. 三维形体的表面积

在 N * N 的网格上，我们放置一些 1 * 1 * 1  的立方体。

每个值 v = grid[i][j] 表示 v 个正方体叠放在对应单元格 (i, j) 上。

请你返回最终形体的表面积。
*/

func surfaceArea(grid [][]int) int {
	n := len(grid)

	cubes := 0
	faces := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				cubes += grid[i][j]
				faces += grid[i][j] - 1
			}
			if i > 0 {
				faces += min(grid[i-1][j], grid[i][j])
			}
			if j > 0 {
				faces += min(grid[i][j-1], grid[i][j])
			}
		}
	}
	return 6*cubes - 2*faces
}
