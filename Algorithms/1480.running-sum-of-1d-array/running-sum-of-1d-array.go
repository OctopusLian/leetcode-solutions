func runningSum(nums []int) []int {
    //中间变量，一边循环，一边保存元素和
    sum := 0
    for i:=0;i<len(nums);i++ {
        //将索引小于等于i的全部元素相加，赋值索引为i的元素
        sum += nums[i]
        nums[i] = sum
    }
    return nums
}