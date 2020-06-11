package src

func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var i = 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
