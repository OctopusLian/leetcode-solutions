func new21Game(N int, K int, W int) float64 {
	//解法1，动态规划
    if K == 0 {
        return 1.0
    }
    dp := make([]float64, K + W)
    for i := K; i <= N && i < K + W; i++ {
        dp[i] = 1.0
    }
    for i := K - 1; i >= 0; i-- {
        for j := 1; j <= W; j++ {
            dp[i] += dp[i + j] / float64(W)
        }
    }
    return dp[0]
}

//超出时间限制