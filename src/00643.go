package src

/*643. 子数组最大平均数 I*/

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) <= k {
		sum := 0
		for i := 0; i < len(nums); i++ {
			sum += nums[i]
		}
		return float64(sum) / float64(len(nums))
	}
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	max := sum
	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}
