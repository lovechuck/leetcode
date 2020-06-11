package src

import "sort"

/*
公司计划面试 2N 人。第 i 人飞往 A 市的费用为 costs[i][0]，飞往 B 市的费用为 costs[i][1]。

返回将每个人都飞到某座城市的最低费用，要求每个城市都有 N 人抵达。
*/
func twoCitySchedCost(costs [][]int) int {
	sum := 0
	more := []int{}
	for i := 0; i < len(costs); i++ {
		sum += costs[i][1]
		more = append(more, costs[i][0]-costs[i][1])
	}

	sort.Ints(more)
	for i := 0; i < len(costs)/2; i++ {
		sum += more[i]
	}
	return sum
}
