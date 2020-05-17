package src

/*
868. 二进制间距

给定一个正整数 N，找到并返回 N 的二进制表示中两个连续的 1 之间的最长距离。

如果没有两个连续的 1，返回 0 。
*/

func binaryGap(N int) int {
	ans := 0
	pre := -1
	i := 0
	for N > 0 {
		if N&1 > 0 {
			if pre != -1 {
				if ans < i-pre {
					ans = i - pre
				}
			}
			pre = i
		}
		i++
		N >>= 1
	}
	return ans
}
