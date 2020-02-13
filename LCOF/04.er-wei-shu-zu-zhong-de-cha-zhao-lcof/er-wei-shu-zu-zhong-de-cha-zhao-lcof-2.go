func findNumberIn2DArray(matrix [][]int, target int) bool {
    for _,nums:=range matrix{
        //遍历数组切片，查找数组中是否含有target值，如果查找不到，返回值是target应该插入数组的位置（会保持数组的递增顺序）
        i:=sort.SearchInts(nums,target)
        //插入的位置小于数组长度 且 插入数组的位置上的值和目标值相等
        if i<len(nums)&&target==nums[i]{
            return true
        }
    }
    return false
}
