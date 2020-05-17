package src

/*
883. 三维形体投影面积

在 N * N 的网格中，我们放置了一些与 x，y，z 三轴对齐的 1 * 1 * 1 立方体。

每个值 v = grid[i][j] 表示 v 个正方体叠放在单元格 (i, j) 上。

现在，我们查看这些立方体在 xy、yz 和 zx 平面上的投影。

投影就像影子，将三维形体映射到一个二维平面上。

在这里，从顶部、前面和侧面看立方体时，我们会看到“影子”。

返回所有三个投影的总面积。
*/
func projectionArea(grid [][]int) int {

	ans := 0
	n := len(grid)
	for i := 0; i < n; i++ {
		xans := 0
		yans := 0
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				ans++
			}
			if xans < grid[i][j] {
				xans = grid[i][j]
			}
			if yans < grid[j][i] {
				yans = grid[j][i]
			}
		}
		ans += xans + yans
	}
	return ans
}
