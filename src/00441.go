package src

/*441. 排列硬币*/
func arrangeCoins(n int) int {
	sum := 0
	i := 1
	for ; i <= n; i++ {
		sum += i
		if sum > n {
			break
		}
	}
	return i - 1
}
