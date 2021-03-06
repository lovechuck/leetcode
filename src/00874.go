package src

/*

874. 模拟行走机器人


机器人在一个无限大小的网格上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令：

-2：向左转 90 度
-1：向右转 90 度
1 <= x <= 9：向前移动 x 个单位长度
在网格上有一些格子被视为障碍物。

第 i 个障碍物位于网格点  (obstacles[i][0], obstacles[i][1])

机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，但仍然可以继续该路线的其余部分。

返回从原点到机器人的最大欧式距离的平方。
*/

func robotSim(commands []int, obstacles [][]int) int {
	type point struct {
		x int
		y int
	}
	obsMap := make(map[point]bool)
	for i := 0; i < len(obstacles); i++ {
		tmp := point{obstacles[i][0], obstacles[i][1]}
		obsMap[tmp] = true
	}
	dir := 0
	x := 0
	y := 0
	res := 0
	points := [][]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}
	for i := 0; i < len(commands); i++ {
		if commands[i] == -1 {
			dir = (dir + 4 - 1) % 4
		} else if commands[i] == -2 {
			dir = (dir + 1) % 4
		} else {
			tmpX := x
			tmpY := y
			for j := 1; j <= commands[i]; j++ {
				if _, ok := obsMap[point{tmpX + points[dir][0], tmpY + points[dir][1]}]; !ok {
					tmpX, tmpY = tmpX+points[dir][0], tmpY+points[dir][1]
				} else {
					x, y = tmpX, tmpY
					break
				}
			}
			x, y = tmpX, tmpY
			if res < x*x+y*y {
				res = x*x + y*y
			}
		}
	}
	return res
}
