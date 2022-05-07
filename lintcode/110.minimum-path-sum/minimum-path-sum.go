/**
 * @param grid: a list of lists of integers
 * @return: An integer, minimizes the sum of all numbers along its path
 */
func minPathSum (grid [][]int) int {
    // write your code here
    //解法1，动态规划
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    rows, columns := len(grid), len(grid[0])
    dp := make([][]int, rows)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, columns)
    }
    dp[0][0] = grid[0][0]
    for i := 1; i < rows; i++ {
        dp[i][0] = dp[i - 1][0] + grid[i][0]
    }
    for j := 1; j < columns; j++ {
        dp[0][j] = dp[0][j - 1] + grid[0][j]
    }
    for i := 1; i < rows; i++ {
        for j := 1; j < columns; j++ {
            dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
        }
    }
    return dp[rows - 1][columns - 1]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
