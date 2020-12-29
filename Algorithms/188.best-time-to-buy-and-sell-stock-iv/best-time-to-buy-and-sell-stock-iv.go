func maxProfit(k int, prices []int) int {
    n := len(prices)
    if n == 0 {
        return 0
    }

    k = min(k, n/2)
    buy := make([][]int, n)
    sell := make([][]int, n)
    for i := range buy {
        buy[i] = make([]int, k+1)
        sell[i] = make([]int, k+1)
    }
    buy[0][0] = -prices[0]
    for i := 1; i <= k; i++ {
        buy[0][i] = math.MinInt64 / 2
        sell[0][i] = math.MinInt64 / 2
    }

    for i := 1; i < n; i++ {
        buy[i][0] = max(buy[i-1][0], sell[i-1][0]-prices[i])
        for j := 1; j <= k; j++ {
            buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
            sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])
        }
    }
    return max(sell[n-1]...)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a ...int) int {
    res := a[0]
    for _, v := range a[1:] {
        if v > res {
            res = v
        }
    }
    return res
}