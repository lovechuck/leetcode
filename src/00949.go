package src

import (
	"fmt"
)

/*
949. 给定数字能组成的最大时间

给定一个由 4 位数字组成的数组，返回可以设置的符合 24 小时制的最大时间。

最小的 24 小时制时间是 00:00，而最大的是 23:59。从 00:00 （午夜）开始算起，过得越久，时间越大。

以长度为 5 的字符串返回答案。如果不能确定有效时间，则返回空字符串。
*/

func largestTimeFromDigits(A []int) string {
	ans := -1
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j != i {
				for k := 0; k < 4; k++ {
					if k != i && k != j {
						g := 6 - i - j - k
						hour := A[i]*10 + A[j]
						minute := A[k]*10 + A[g]
						if hour < 24 && minute < 60 {
							if ans < hour*60+minute {
								ans = hour*60 + minute
							}
						}
					}
				}
			}
		}
	}
	if ans > -1 {
		return fmt.Sprintf("%02d:%02d", ans/60, ans%60)
	}
	return ""
}

func Test_largestTimeFromDigits() {
	arr := []int{1, 2, 3, 4}
	largestTimeFromDigits(arr)
}
