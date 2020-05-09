package src

/*
661. 图片平滑器

包含整数的二维矩阵 M 表示一个图片的灰度。你需要设计一个平滑器来让每一个单元的灰度成为平均灰度 (向下舍入) ，平均灰度的计算是周围的8个单元和它本身的值求平均，如果周围的单元格不足八个，则尽可能多的利用它们。
*/
func imageSmoother(M [][]int) [][]int {
	x := len(M)
	y := len(M[0])

	res := make([][]int, x)
	for i := 0; i < x; i++ {
		tmp := make([]int, y)
		for j := 0; j < y; j++ {
			sum := 0
			count := 0
			for xr := i - 1; xr <= i+1; xr++ {
				for yr := j - 1; yr <= j+1; yr++ {
					if xr >= 0 && xr < x && yr >= 0 && yr < y {
						sum += M[xr][yr]
						count++
					}
				}
			}
			avg := sum / count
			tmp[j] = avg
		}
		res[i] = tmp
	}
	return res
}
