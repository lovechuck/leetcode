package src

import "sort"

/*561. 数组拆分 I*/
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}

	return sum
}
