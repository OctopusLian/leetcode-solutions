func sumOfDistancesInTree(N int, edges [][]int) []int {
	//解法1，树形动态规划
    graph := make([][]int, N)
    for _, e := range edges {
        u, v := e[0], e[1]
        graph[u] = append(graph[u], v)
        graph[v] = append(graph[v], u)
    }

    sz := make([]int, N)
    dp := make([]int, N)
    var dfs func(u, f int)
    dfs = func(u, f int) {
        sz[u] = 1
        for _, v := range graph[u] {
            if v == f {
                continue
            }
            dfs(v, u)
            dp[u] += dp[v] + sz[v]
            sz[u] += sz[v]
        }
    }
    dfs(0, -1)

    ans := make([]int, N)
    var dfs2 func(u, f int)
    dfs2 = func(u, f int) {
        ans[u] = dp[u]
        for _, v := range graph[u] {
            if v == f {
                continue
            }
            pu, pv := dp[u], dp[v]
            su, sv := sz[u], sz[v]

            dp[u] -= dp[v] + sz[v]
            sz[u] -= sz[v]
            dp[v] += dp[u] + sz[u]
            sz[v] += sz[u]

            dfs2(v, u)

            dp[u], dp[v] = pu, pv
            sz[u], sz[v] = su, sv
        }
    }
    dfs2(0, -1)
    return ans
}