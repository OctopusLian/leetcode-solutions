func findLengthOfLCIS(nums []int) (ans int) {
	//解法1，贪心
    start := 0
    for i, v := range nums {
        if i > 0 && v <= nums[i-1] {
            start = i
        }
        ans = max(ans, i-start+1)
    }
    return
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}