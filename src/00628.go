package src

import (
	"sort"
)

/*
628. 三个数的最大乘积
*/

func maximumProduct(nums []int) int {
	leng := len(nums)
	sort.Ints(nums)
	last := nums[leng-3] * nums[leng-2] * nums[leng-1]
	start := nums[0] * nums[1] * nums[leng-1]
	if last > start {
		return last
	}
	return start
}
