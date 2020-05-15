func waysToStep(n int) int {
	//解法2，动态规划
    dp := make([]int, n+4)
    dp[0], dp[1], dp[2], dp[3] = 0, 1, 2, 4
    for i := 4;i <= n;i++ {
        //dp[i] = dp[i-3] + dp[i-2] + dp[i-3]
        dp[i] = (dp[i-1] % 1000000007 + dp[i-2] % 1000000007 + dp[i-3] % 1000000007) % 1000000007
    } 
    return dp[n] % 1000000007
}