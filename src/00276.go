package src

/*
 * 动态规划，假设连续三个栅栏分别为1，2，3；
 * 则栅栏3要么和1，要么和2颜色不同,当前位置有k-1种可能颜色；
 * 因此dp3 = (k-1) * dp1 + (k-1) * dp2;
 *
 */

func numWays(n int, k int) int {
	dp := make([]int, n)
	dp[0] = k
	dp[1] = k * (k - 1)
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1]*(k-1) + dp[k-2]*(k-1)
	}
	return dp[n-1]
}
