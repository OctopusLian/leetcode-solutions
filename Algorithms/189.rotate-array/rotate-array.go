func rotate(nums []int, k int)  {
    //暴力解决
    var temp,previous int
    for i:=0;i<k;i++ {
        previous = nums[len(nums)-1]
        for j:=0;j<len(nums);j++ {
            temp = nums[j]
            nums[j] = previous
            previous = temp
        }
    }
}
