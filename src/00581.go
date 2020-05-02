package src

import "fmt"

/*581. 最短无序连续子数组*/

func findUnsortedSubarray(nums []int) int {

	if nums == nil {
		return 0
	}
	var start []int
	begin := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= nums[i+1] {
			start = append(start, i)
		} else {
			start = append(start, i)
			begin = i
			break
		}
	}
	if begin == -1 {
		return 0
	}
	var finish []int
	end := len(nums) - 1
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] >= nums[i-1] {
			finish = append(finish, i)
		} else {
			finish = append(finish, i)
			end = i
			break
		}
	}

	min, max := nums[begin], nums[end]
	for i := begin; i <= end; i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	s, d := -1, -1
	for i := 0; i < len(start); i++ {
		if min < nums[start[i]] {
			s = start[i]
			break
		}
	}
	for j := 0; j < len(finish); j++ {
		// 从后往前
		if max > nums[finish[j]] {
			d = finish[j]
			break
		}
	}

	return d - s + 1
}

func Test_findUnsortedSubarray() {
	nums := []int{1, 3, 2, 2, 2}
	res := findUnsortedSubarray(nums)
	fmt.Println(res)
}
