func missingNumber(nums []int) int {
    nummap := make(map[int]bool)
    //将nums数组里的每个数字都存进map里
    for _,num := range nums {
        nummap[num] = true
    }
    //从0开始，在map里遍历，如果没找到，那就是缺失的数字，否则返回-1
    for i:=0;i<len(nums)+1;i++ {
        if !(nummap[i]) {
            return i
        }
    }

    return -1
}
