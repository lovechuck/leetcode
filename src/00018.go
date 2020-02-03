package src

import (
	"fmt"
	"sort"
)

/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

func judgeExist(list [][]int, temp []int) bool {
	var time int
	for i := 0; i < len(list); i++ {
		time = 0
		for j := 0; j < len(temp); j++ {
			if list[i][j] != temp[j] {
				break
			} else {

				time++
			}
			if time == len(temp) {

				return true
			}
		}
	}
	return false
}

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length-3; i++ {
		l := i + 1
		m := l + 1
		for l < length-2 {
			r := length - 1
			for m < length-1 && m < r {
				var tmpArr []int
				sum := nums[i] + nums[l] + nums[m] + nums[r]
				if sum == target {
					tmpArr = append(tmpArr, nums[i])
					tmpArr = append(tmpArr, nums[l])
					tmpArr = append(tmpArr, nums[m])
					tmpArr = append(tmpArr, nums[r])

					if !judgeExist(result, tmpArr) {
						result = append(result, tmpArr)
					}
					m++
					r--

					for m < length-1 && nums[m] == nums[m-1] {
						m++
					}
					for r > 0 && nums[r] == nums[r+1] {
						r--
					}
				} else if sum < target {
					m++
				} else if sum > target {
					r--
				}
			}
			l++
			m = l + 1
		}
	}
	return result
}

func Test_fourSum() {
	nums := []int{0, 0, 0, 0}
	target := 0
	result := fourSum(nums, target)
	fmt.Println(result)
}
