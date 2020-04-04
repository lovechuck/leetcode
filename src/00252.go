package src

import "sort"

/*
给定一个会议时间安排的数组，每个会议时间都会包括开始和结束的时间 [[s1,e1],[s2,e2],...] (si < ei)，请你判断一个人是否能够参加这里面的全部会议。

示例 1:

输入:
[[0,30],[5,10],[15,20]]

输出: false
示例 2:

输入: [[7,10],[2,4]]
输出: true
*/

func canAttendMeetings(rooms [][2]int) bool {
	sort.Slice(rooms, func(i, j int) bool {
		return rooms[i][0] < rooms[j][0]
	})

	for j = 1; j < len(rooms); j++ {
		if rooms[j-1][1] > rooms[j][0] {
			return false
		}
	}

	return true
}
