package src

import "math"

/*
849. 到最近的人的最大距离

输入：[1,0,0,0,1,0,1]
输出：2
解释：
如果亚历克斯坐在第二个空位（seats[2]）上，他到离他最近的人的距离为 2 。
如果亚历克斯坐在其它任何一个空位上，他到离他最近的人的距离为 1 。
因此，他到离他最近的人的最大距离是 2 。
*/

func maxDistToClosest(seats []int) int {
	if seats == nil {
		return 0
	}
	pre := -1
	count := make([]int, len(seats))
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			pre = i
		}
		if pre != -1 {
			count[i] = i - pre
		} else {
			count[i] = math.MaxInt32
		}
	}
	pre = -1
	for i := len(seats) - 1; i >= 0; i-- {
		if seats[i] == 1 {
			pre = i
		}
		if pre != -1 {
			if pre-i < count[i] {
				count[i] = pre - i
			}
		}
	}
	max := count[0]
	for _, i := range count {
		if max < i {
			max = i
		}
	}
	return max
}

func Test_maxDistToClosest() {
	maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1})
}
