package src

/*35. 搜索插入位置*/

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			return mid
		} else {
			left = mid + 1
		}
	}

	return right
}
