package src

/*
题目描述
假如有一排房子，共 n 个，每个房子可以被粉刷成红色、蓝色或者绿色这三种颜色中的一种，你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。

当然，因为市场上不同颜色油漆的价格不同，所以房子粉刷成不同颜色的花费成本也是不同的。每个房子粉刷成不同颜色的花费是以一个 n x 3 的矩阵来表示的。

例如，costs[0][0] 表示第 0 号房子粉刷成红色的成本花费；costs[1][2] 表示第 1 号房子粉刷成绿色的花费，以此类推。请你计算出粉刷完所有房子最少的花费成本。

注意：

所有花费均为正整数。

示例：

输入: [[17,2,17],[16,16,5],[14,3,19]]
输出: 10
解释: 将 0 号房子粉刷成蓝色，1 号房子粉刷成绿色，2 号房子粉刷成蓝色。
最少花费: 2 + 5 + 3 = 10。
*/

func minCost(color [][3]int) int {
	n := len(color)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return min3(color[0][0], color[0][1], color[0][2])
	}
	cost := make([][3]int, n)
	cost[0][0] = color[0][0]
	cost[0][1] = color[0][1]
	cost[0][2] = color[0][2]
	for i := 1; i < n; i++ {
		cost[i][0] = color[i][0] + min(cost[i-1][1], cost[i-1][2])
		cost[i][1] = color[i][1] + min(cost[i-1][0], cost[i-1][2])
		cost[i][2] = color[i][2] + min(cost[i-1][0], cost[i-1][1])
	}

	return min3(cost[n-1][0], cost[n-1][1], cost[n-1][2])
}

func min3(a int, b int, c int) int {
	min := 0
	if a < b {
		min = a
	} else {
		min = b
	}
	if min < c {
		return min
	} else {
		return c
	}
}
