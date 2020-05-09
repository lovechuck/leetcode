package src

/*665. 非递减数列*/
func checkPossibility(nums []int) bool {
	p := 0
	index := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			p++
			index = i
		}
		if p > 1 {
			return false
		}
	}
	if p == 0 {
		return true
	}
	if index == 0 || index == len(nums)-2 {
		return true
	}
	if nums[index-1] <= nums[index+1] || nums[index] <= nums[index+2] {
		return true
	}
	return false
}
