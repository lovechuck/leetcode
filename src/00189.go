package src

func rotate(nums []int, k int) {
	if nums == nil || len(nums) == 1 {
		return
	}
	k %= len(nums)
	reverse_r(nums, 0, len(nums)-1)
	reverse_r(nums, 0, k-1)
	reverse_r(nums, k, len(nums)-1)
}

func reverse_r(nums []int, start int, end int) {
	var temp int
	for start < end {
		temp = nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}
