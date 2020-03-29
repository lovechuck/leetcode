package src

/*198. 打家劫舍*/

func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return mmax(nums[0], nums[1])
	}
	result := make([]int, len(nums))
	result[0] = nums[0]
	result[1] = mmax(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		result[i] = mmax(result[i-1], result[i-2]+nums[i])
	}
	return result[len(nums)-1]
}
