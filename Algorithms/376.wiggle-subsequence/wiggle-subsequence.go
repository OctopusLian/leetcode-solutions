func wiggleMaxLength(nums []int) int {
    n := len(nums)
    if n < 2 {
        return n
    }
    up := make([]int, n)
    down := make([]int, n)
    up[0] = 1
    down[0] = 1
    for i := 1; i < n; i++ {
        if nums[i] > nums[i-1] {
            up[i] = max(up[i-1], down[i-1]+1)
            down[i] = down[i-1]
        } else if nums[i] < nums[i-1] {
            up[i] = up[i-1]
            down[i] = max(up[i-1]+1, down[i-1])
        } else {
            up[i] = up[i-1]
            down[i] = down[i-1]
        }
    }
    return max(up[n-1], down[n-1])
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}