func subsets(nums []int) [][]int {
    var ans [][]int
    len := len(nums)
    count := 1 << len // 获取2^n这个个数常量
    for n := 0; n < count; n++ { // 外循环对应产生的数字
        ans = append(ans, []int{})
        for mask, i := 1, len-1; mask < count; mask, i = mask<<1, i-1 { // 里循环对应集合长度 并按位取值
            if (n & mask) != 0 { // 如果非0 那么表示这位上有 然后加之
                ans[n] = append(ans[n], nums[i])
            }
        }
    }
    return ans
}