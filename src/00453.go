package src

import "math"

/*
453. 最小移动次数使数组元素相等
给定一个长度为 n 的非空整数数组，找到让数组所有元素相等的最小移动次数。每次移动可以使 n - 1 个元素增加 1。

示例:

输入:
[1,2,3]

输出:
3
*/

func minMoves(nums []int) int {
	// 找到最小的
	min := math.MaxInt32
	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	sum := 0
	//每次移动到和最大的相等
	for _, num := range nums {
		sum += num - min
	}
	return sum
}
