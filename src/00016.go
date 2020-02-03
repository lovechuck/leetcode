package src

import (
	"sort"
)

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
*/

func threeSumClosest(nums []int, target int) int {
	if len(nums) <= 2 {
		return 0
	}
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}

	result := nums[0] + nums[1] + nums[2]
	prev := absInt(result - target)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if absInt(sum-target) < prev {
				prev = absInt(sum - target)
				result = sum
			}
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				result = target
				break
			}
		}
	}

	return result
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
