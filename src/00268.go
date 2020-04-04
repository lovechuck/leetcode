package src

/*268. 缺失数字*/
func missingNumber(nums []int) int {
	in := len(nums)
	sum := in * (in + 1) / 2
	total := 0
	for i := 0; i < in; i++ {
		total = total + nums[i]
	}

	return sum - total
}
