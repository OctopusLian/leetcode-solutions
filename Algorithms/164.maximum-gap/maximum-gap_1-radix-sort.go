func maximumGap(nums []int) (ans int) {
	//解法1，基数排序
    n := len(nums)
    if n < 2 {
        return
    }

    buf := make([]int, n)
    maxVal := max(nums...)
    for exp := 1; exp <= maxVal; exp *= 10 {
        cnt := [10]int{}
        for _, v := range nums {
            digit := v / exp % 10
            cnt[digit]++
        }
        for i := 1; i < 10; i++ {
            cnt[i] += cnt[i-1]
        }
        for i := n - 1; i >= 0; i-- {
            digit := nums[i] / exp % 10
            buf[cnt[digit]-1] = nums[i]
            cnt[digit]--
        }
        copy(nums, buf)
    }

    for i := 1; i < n; i++ {
        ans = max(ans, nums[i]-nums[i-1])
    }
    return
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