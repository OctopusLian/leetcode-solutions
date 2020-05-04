func missingNumber(nums []int) int {
	//1，新建一个`map`，使用`for`循环记录`nums`数组中出现的元素
    var numsmap = make(map[int]bool)
    for _,n := range nums {
        numsmap[n] = true
	}
	//使用`for`循环从零开始累加，在`map`中查找这个元素是否存在，如果不存在，返回这个元素，否则返回-1
    for i:=0;i<=len(nums);i++ {
        if !numsmap[i] {
            return i
        }
    }

    return -1
}