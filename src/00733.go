package src

/*733. 图像渲染*/

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	if color != newColor {
		floodFill_dfs(image, sr, sc, color, newColor)
	}
	return image
}

func floodFill_dfs(image [][]int, sr int, sc int, color int, newColor int) {
	if image[sr][sc] == color {
		image[sr][sc] = newColor
		if sr-1 >= 0 {
			floodFill_dfs(image, sr-1, sc, color, newColor)
		}
		if sc-1 >= 0 {
			floodFill_dfs(image, sr, sc-1, color, newColor)
		}
		if sr+1 < len(image) {
			floodFill_dfs(image, sr+1, sc, color, newColor)
		}
		if sc+1 < len(image[0]) {
			floodFill_dfs(image, sr, sc+1, color, newColor)
		}
	}
}
