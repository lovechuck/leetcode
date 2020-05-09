package src

/*674. 最长连续递增序列*/

func findLengthOfLCIS(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	ans, start := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			start = i + 1
		}
		if ans < i+1-start+1 {
			ans = i + 1 - start + 1
		}
	}
	return ans
}
