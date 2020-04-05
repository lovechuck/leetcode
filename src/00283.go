package src

/*283. 移动零*/

func moveZeroes(nums []int) {
	last := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[last] = nums[i]
			last++
		}
	}
	for last < len(nums) {
		nums[last] = 0
		last++
	}
}
