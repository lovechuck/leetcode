package src

import "fmt"

/*
 二进制手表
*/
func readBinaryWatch(num int) []string {
	var res []string
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if count(i)+count(j) == num {
				if j < 10 {
					res = append(res, fmt.Sprintf("%d:0%d", i, j))
				} else {
					res = append(res, fmt.Sprintf("%d:%d", i, j))
				}
			}
		}
	}
	return res
}

//计算二进制中1的个数
func count(n int) int {
	res := 0
	for n != 0 {
		n = n & (n - 1)
		res++
	}

	return res
}
