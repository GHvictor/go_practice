package leetcode

/**
DP
 */
func numTrees(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j]*dp[i-1-j]
		}
	}
	return dp[n]
}
