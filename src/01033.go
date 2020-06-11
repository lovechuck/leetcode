package src

import "sort"

/*
1033. 移动石子直到连续
*/
func numMovesStones(a int, b int, c int) []int {
	nums := []int{a, b, c}
	sort.Ints(nums)
	max := nums[2] - nums[0] - 2
	if nums[1]-nums[0] == 1 && nums[2]-nums[1] == 1 {
		return []int{0, max}
	}
	if nums[1]-nums[0] <= 2 || nums[2]-nums[1] <= 2 {
		return []int{1, max}
	}
	return []int{2, max}
}
