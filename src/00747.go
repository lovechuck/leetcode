package src

import "math"

/*
747. 至少是其他数字两倍的最大数
*/
func dominantIndex(nums []int) int {
	if nums == nil {
		return -1
	}
	max := math.MinInt32
	maxIndex := 0
	for i, num := range nums {
		if num >= max {
			maxIndex = i
			max = num
		}
	}
	for i, num := range nums {
		if i != maxIndex && 2*num > max {
			return -1
		}
	}

	return maxIndex
}
