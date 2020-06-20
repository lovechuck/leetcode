package src

import "sort"

/*
1051. 高度检查器


学校在拍年度纪念照时，一般要求学生按照 非递减 的高度顺序排列。

请你返回能让所有学生以 非递减 高度排列的最小必要移动人数。

注意，当一组学生被选中时，他们之间可以以任何可能的方式重新排序，而未被选中的学生应该保持不动。
*/

func heightChecker(heights []int) int {
	var old []int
	for i := 0; i < len(heights); i++ {
		old = append(old, heights[i])
	}
	sort.Ints(heights)

	ans := 0
	for i := 0; i < len(heights); i++ {
		if old[i] != heights[i] {
			ans++
		}
	}

	return ans
}
