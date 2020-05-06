func mincostTickets(days []int, costs []int) int {
	//记忆化搜索（窗口变量型）
	//dp(i)=min(dp(j1​)+costs[0],dp(j7​)+costs[1],dp(j30​)+costs[2])
    memo := [366]int{}
    durations := []int{1, 7, 30}

    var dp func(idx int) int 
    dp = func(idx int) int {
        if idx >= len(days) {
            return 0
        }
        if memo[idx] > 0 {
            return memo[idx]
        }
        memo[idx] = math.MaxInt32
        j := idx
        for i := 0; i < 3; i++ {
            for ; j < len(days) && days[j] < days[idx] + durations[i]; j++ { }
            memo[idx] = min(memo[idx], dp(j) + costs[i])
        }
        return memo[idx]
    }
    return dp(0)
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}