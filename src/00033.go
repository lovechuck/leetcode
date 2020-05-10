package src

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:

输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
示例 2:

输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
*/

func find_rotate_index(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return 0
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] {
			return mid + 1
		}
		if nums[mid-1] > nums[mid] {
			return mid
		}
		if nums[mid] > nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return 0
}

func find(nums []int, target int, left int, right int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func search2(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	rotate_index := find_rotate_index(nums)
	if nums[rotate_index] == target {
		return rotate_index
	}
	if rotate_index == 0 {
		return find(nums, target, 0, len(nums)-1)
	}

	if nums[0] > target {
		return find(nums, target, rotate_index, len(nums)-1)
	}
	return find(nums, target, 0, rotate_index)
}
