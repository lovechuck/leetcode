package src

func maxProfit2(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	length := len(prices)
	dp := make([][2]int, length)
	for i := 0; i < length; i++ {
		if i-1 < 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
		} else {
			dp[i][0] = mmax(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = mmax(dp[i-1][1], dp[i-1][0]-prices[i])
		}
	}
	return dp[length-1][0]
}
