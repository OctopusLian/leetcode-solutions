func numberOfSubarrays(nums []int, k int) int {
    //方法1，动态规划
    dp := make([]int, 0)
	cnt, ret := 0, 0
	for i := 0; i < len(nums); i++ {
		cnt++
		if nums[i]%2 == 1 {
			dp = append(dp, cnt)
			cnt = 0
		}
		if len(dp) >= k {
			ret += dp[len(dp) - k]
		}
	}
	return ret
}
