func maxSubArray(nums []int) int {
    //一个临时值，和一个最大值
    temp,max := 0,nums[0]
    for i:=0;i<len(nums);i++ {
        if (temp + nums[i]) > nums[i] {
            temp = temp + nums[i]
        } else {
            temp = nums[i]
        }
        if temp > max {
            max = temp
        }
    }

    return max  //返回最大值
}
