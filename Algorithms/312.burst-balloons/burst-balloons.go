func maxCoins(nums []int) int {
	//解法1，动态规划
    n := len(nums)
    rec := make([][]int, n + 2)
    for i := 0; i < n + 2; i++ {
        rec[i] = make([]int, n + 2)
    }
    val := make([]int, n + 2)
    val[0], val[n+1] = 1, 1
    for i := 1; i <= n; i++ {
        val[i] = nums[i-1]
    }
    for i := n - 1; i >= 0; i-- {
        for j := i + 2; j <= n + 1; j++ {
            for k := i + 1; k < j; k++ {
                sum := val[i] * val[k] * val[j]
                sum += rec[i][k] + rec[k][j]
                rec[i][j] = max(rec[i][j], sum)
            }
        }
    }
    return rec[0][n+1]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}