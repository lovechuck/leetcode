package src

/*
1030. 距离顺序排列矩阵单元格
*/
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {

	ans := make([][][]int, R*C)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			dist := distance(i, j, r0, c0)
			ans[dist] = append(ans[dist], []int{i, j})
		}
	}

	var res [][]int
	for _, an := range ans {
		if an != nil {
			res = append(res, an...)
		}
	}
	return res
}
func distance(a int, b int, a1 int, b1 int) int {
	t1 := 0
	t2 := 0
	if a > a1 {
		t1 = a - a1
	} else {
		t1 = a1 - a
	}
	if b > b1 {
		t2 = b - b1
	} else {
		t2 = b1 - b
	}
	return t1 + t2
}
