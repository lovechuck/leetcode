package src

import "math"

/*给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。*/

func thirdMax(nums []int) int {
	first, second, third := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num > third { // 通过第3关
			if num < second {
				third = num
			} else if num > second { // 通过第2关
				if num < first {
					third = second
					second = num
				} else if num > first { // 通过第1关
					third = second
					second = first
					first = num
				}
			}
		}
	}
	if third != math.MinInt64 {
		return third
	} else {
		return first
	}
}
