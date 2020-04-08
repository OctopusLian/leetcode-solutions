func movingCount(m int, n int, k int) int {
    // 使用匿名函数 - 需要提前声明才能调用自身
	var dfs func(y, x int) (res int)

	// 用Map存储已访问过的坐标 - 利用Map的特性
	visited := make(map[string]bool)

	// 计算数位和
	digitSum := func(x int) (res int) {
		for ; x > 0; x = x / 10 {
			res += x % 10
		}
		return
	}

	// 深度优先搜索
	dfs = func(y, x int) (res int) {
		// 判断坐标是否越界
		if y < 0 || y > m-1 {
			return
		}
		if x < 0 || x > n-1 {
			return
		}
		
		// 判断坐标是否已经访问过
		if _, isExist := visited[fmt.Sprintf("%d,%d", y, x)]; isExist {
			return
		}
		
		// 判断数位和是否大于k
		if (digitSum(y) + digitSum(x)) > k {
			return
		}
		
		// 将该坐标加入已访问
		visited[fmt.Sprintf("%d,%d", y, x)] = true
		res += 1
		
		// 搜索相邻的坐标
		res += dfs(y+1, x)
		res += dfs(y, x+1)
		return
	}

    // 执行DFS算法
    return dfs(0, 0)
}
