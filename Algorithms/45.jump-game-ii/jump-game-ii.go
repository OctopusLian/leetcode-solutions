func jump(nums []int) int {
    // 动态规划四步走：
    // 1、状态：f[i] 表示从起点到当前位置跳的最小次数
    // 2、推导：f[i] = min(f[j]+1),a[j]+j >=i 表示从j位置用一步跳到当前位置，这个j位置可能有多个，取最小的一个就行
    // 3、初始化：f[0] = 0
    // 4、结果：f[n-1]
    f := make([]int, len(nums))
    f[0] = 0
    for i := 1; i < len(nums); i++ {
        // f[i] 先取一个默认最大值i
        f[i] = i
        // 遍历之前结果取一个最小值+1
        for j := 0; j < i; j++ {
            if nums[j]+j >= i {
                f[i] = min(f[j]+1,f[i])
            }
        }
    }
    return f[len(nums)-1]
}
func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
