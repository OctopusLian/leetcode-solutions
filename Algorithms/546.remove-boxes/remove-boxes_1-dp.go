func removeBoxes(boxes []int) int {
	//解法1，动态规划
    dp := [100][100][100]int{}
    var calculatePoints func(boxes []int, l, r, k int) int
    calculatePoints = func(boxes []int, l, r, k int) int {
        if l > r {
            return 0
        }
        if dp[l][r][k] != 0 {
            return dp[l][r][k]
        }
        for r > l && boxes[r] == boxes[r - 1] {
            r--
            k++
        }
        dp[l][r][k] = calculatePoints(boxes, l, r - 1, 0) + (k + 1) * (k + 1)
        for i := l; i < r; i++ {
            if boxes[i] == boxes[r] {
                dp[l][r][k] = max(dp[l][r][k], calculatePoints(boxes, l, i, k + 1) + calculatePoints(boxes, i + 1, r - 1, 0))
            }
        }
        return dp[l][r][k]
    }
    return calculatePoints(boxes, 0, len(boxes) - 1, 0)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}