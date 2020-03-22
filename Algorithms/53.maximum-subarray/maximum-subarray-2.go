func maxSubArray(nums []int) int {
    //dp解法
    var dp = map[int]int{}
    dp[0] = nums[0]
    max := nums[0]
    for i := 1; i < len(nums); i++ {
        //做判断
        if (dp[i - 1] + nums[i]) > nums[i] {
            dp[i] = dp[i - 1] + nums[i]
        } else {
            dp[i] = nums[i]
        }

        if dp[i] > max {
            max = dp[i]
        }
    }
    return max
}
