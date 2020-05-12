package src

/*

746. 使用最小花费爬楼梯
cost[i] = min(cost[i-2],cost[i-1])+xi
*/
func minCostClimbingStairs(cost []int) int {
	if len(cost) <= 2 {
		return 0
	}
	var fn []int
	for i := 0; i < len(cost); i++ {
		if i == 0 {
			fn = append(fn, cost[i])
		} else if i == 1 {
			fn = append(fn, cost[i])
		} else if fn[i-1] < fn[i-2] {
			fn = append(fn, fn[i-1]+cost[i])
		} else {
			fn = append(fn, fn[i-2]+cost[i])
		}
	}
	a := len(fn)
	if fn[a-1] < fn[a-2] {
		return fn[a-1]
	} else {
		return fn[a-2]
	}
}
