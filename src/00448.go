package src

/*448. 找到所有数组中消失的数字*/

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		temp := abs_nums(nums[i]) - 1
		if nums[temp] > 0 {
			nums[temp] = -nums[temp]
		}
	}
	var res []int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs_nums(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
