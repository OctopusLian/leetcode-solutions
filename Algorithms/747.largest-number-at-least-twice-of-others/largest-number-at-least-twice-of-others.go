func dominantIndex(nums []int) int {
    //小于等于一个元素 直接输出
    if len(nums)<=1{
        return 0
    }
    // 初始化 次大为0 最大为1
    lastMax,max:=0,nums[0]
    // 最大的下标
    maxIndex:=0
    // 遍历
    for i:=1;i<len(nums);i++{
        // 如果num[i]比max大，max变为次大，num[i]变为最大
        if nums[i]>max{
            lastMax=max
            max=nums[i]
            maxIndex=i
        // 如果比lastMax大 替换
        }else if nums[i]>lastMax{
            lastMax=nums[i]
        }
    }
    // 判断
    if max>=lastMax*2{
        return maxIndex
    }

    return -1
}