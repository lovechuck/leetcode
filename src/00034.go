package src

import "fmt"

/*
34. 在排序数组中查找元素的第一个和最后一个位置
*/

func findLeft(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if left <= len(nums)-1 && nums[left] == target {
		return left
	}

	return -1
}

func findRight(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			left = mid + 1
		} else {
			left = mid + 1
		}
	}

	if right >= 0 && nums[right] == target {
		return right
	}

	return -1
}

func searchRange(nums []int, target int) []int {
	left := findLeft(nums, target)
	right := findRight(nums, target)
	return []int{left, right}
}

func Test_searchRange() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	arr := searchRange(nums, target)
	fmt.Println(arr)
}
