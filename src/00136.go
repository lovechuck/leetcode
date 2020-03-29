package src

/*
136. 只出现一次的数字
*/

func singleNumber(nums []int) int {
	a := 0
	for i := 0; i < len(nums); i++ {
		a ^= nums[i]
	}

	return a
}
