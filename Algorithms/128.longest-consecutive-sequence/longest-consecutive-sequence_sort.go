func longestConsecutive(nums []int) int {
    n := len(nums)
    if n <= 1 {
        return n
    }
    sort.Ints(nums) // 排序
    ret, tmp := 1, 1
    for k, v := range nums {
        if k <= n-2 {
            if v == nums[k+1] { // 去重
                continue
            }
            if v+1 == nums[k+1] { // 连续
                tmp++
                continue
            }
        }
        if ret < tmp { // 保存最大值
            ret = tmp
        }
        tmp = 1 // 初期话缓存
    }
    return ret
}