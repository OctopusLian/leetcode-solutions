func maxSubArray(nums []int) int {
    b:=nums[0]
    max:=b
    for i:=1;i<len(nums);i++{
        if b<0{
            b=nums[i]
        }else{
            b+=nums[i]
        }
        if b>max{
            max=b
        }
    }
    return max
}