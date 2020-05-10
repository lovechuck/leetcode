package src

import "math"

/*697. 数组的度*/

func findShortestSubArray(nums []int) int {
	if nums == nil {
		return 0
	}
	type obj struct {
		Count int
		Left  int
		Right int
	}
	max := 1
	min := math.MaxInt32
	hash := make(map[int]obj)
	for i := 0; i < len(nums); i++ {
		if v, ok := hash[nums[i]]; ok {
			hash[nums[i]] = obj{
				Count: v.Count + 1,
				Left:  v.Left,
				Right: i,
			}
			if v.Count+1 > max {
				max = v.Count + 1
				min = i - v.Left + 1
			} else if v.Count+1 == max {
				if min > i-v.Left+1 {
					min = i - v.Left + 1
				}
			}
		} else {
			hash[nums[i]] = obj{
				Count: 1,
				Left:  i,
				Right: 0,
			}
		}
	}
	if min == math.MaxInt32 {
		return 1
	}
	return min
}
