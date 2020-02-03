package src

import (
	"fmt"
	"sort"
)

/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

func threeSum(nums []int) [][]int {
	mmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mmap[nums[i]] = i
	}
	exits := make(map[string]bool)
	var values [][]int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			a := nums[i] + nums[j]
			b := -a
			if _, ok := mmap[b]; ok {
				if mmap[b] != i && mmap[b] != j {
					tmp := []int{nums[i], nums[j], b}
					sort.Ints(tmp)
					key := fmt.Sprintf("%d%d%d", tmp[0], tmp[1], tmp[2])
					if _, ok := exits[key]; !ok {
						exits[key] = true
						values = append(values, tmp)
					}
				}
			}
		}
	}

	return values
}
