package src

/*
岛屿的周长
*/
func islandPerimeter(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				//	上  右 下 左
				if i-1 < 0 || grid[i-1][j] == 0 {
					sum += 1
				}
				if j+1 >= len(grid[i]) || grid[i][j+1] == 0 {
					sum += 1
				}
				if i+1 >= len(grid) || grid[i+1][j] == 0 {
					sum += 1
				}
				if j-1 < 0 || grid[i][j-1] == 0 {
					sum += 1
				}
			}
		}
	}
	return sum
}
