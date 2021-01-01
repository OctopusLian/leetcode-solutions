func eraseOverlapIntervals(intervals [][]int) int {
    n := len(intervals)
    if n == 0 {
        return 0
    }
    sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
    dp := make([]int, n)
    for i := range dp {
        dp[i] = 1
    }
    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if intervals[j][1] <= intervals[i][0] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
    }
    return n - max(dp...)
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