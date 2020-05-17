package src

import "math"

/*
821. 字符的最短距离
*/

func shortestToChar(S string, C byte) []int {
	pre := -1
	ans := make([]int, len(S))
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			pre = i
		}
		if pre != -1 {
			ans[i] = i - pre
		} else {
			ans[i] = math.MaxInt32
		}
	}
	pre = -1
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == C {
			pre = i
		}
		if pre != -1 {
			if ans[i] > pre-i {
				ans[i] = pre - i
			}
		}
	}

	return ans
}
