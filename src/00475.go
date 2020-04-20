package src

import (
	"sort"
)

/*
475. 供暖器

冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。

现在，给出位于一条水平线上的房屋和供暖器的位置，找到可以覆盖所有房屋的最小加热半径。

所以，你的输入将会是房屋和供暖器的位置。你将输出供暖器的最小加热半径。

说明:

给出的房屋和供暖器的数目是非负数且不会超过 25000。
给出的房屋和供暖器的位置均是非负数且不会超过10^9。
只要房屋位于供暖器的半径内(包括在边缘上)，它就可以得到供暖。
所有供暖器都遵循你的半径标准，加热的半径也一样。
*/

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	res := 0
	for _, house := range houses {
		cur := indexOf(heaters, house)
		if res < cur {
			res = cur
		}
	}
	return res
}

func indexOf(heaters []int, house int) int {
	left := 0
	right := len(heaters) - 1
	for left <= right {
		mid := (left + right) / 2
		if heaters[mid] == house {
			return 0
		} else if heaters[mid] > house {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	//查找结果为与 right left 之间
	if right < 0 {
		return heaters[left] - house
	} else if left > len(heaters)-1 {
		return house - heaters[right]
	} else {
		return min(heaters[left]-house, house-heaters[right])
	}
}
