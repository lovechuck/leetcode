package src

/*461. 汉明距离*/
func hammingDistance(x int, y int) int {
	xor := x ^ y
	dis := 0
	for xor > 0 {
		dis += 1
		xor = xor & (xor - 1)
	}
	return dis
}
