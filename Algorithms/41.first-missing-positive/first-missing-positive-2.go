func firstMissingPositive(nums []int) int {
	//解法2，哈希表
    n := len(nums)
    for i := 0; i < n; i++ {
        if nums[i] <= 0 {
            nums[i] = n + 1
        }
    }
    for i := 0; i < n; i++ {
        num := abs(nums[i])
        if num <= n {
            fmt.Println(num-1)
            nums[num - 1] = -abs(nums[num - 1])
        }
    }
    for i := 0; i < n; i++ {
        if nums[i] > 0 {
            return i + 1
        }
    }
    return n + 1
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}