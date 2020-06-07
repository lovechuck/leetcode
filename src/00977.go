package src

import "sort"

/*
977. 有序数组的平方

给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。
*/
func sortedSquares(A []int) []int {
	var a []int
	for i := 0; i < len(A); i++ {
		a = append(a, A[i]*A[i])
	}
	sort.Ints(a)
	return a
}
