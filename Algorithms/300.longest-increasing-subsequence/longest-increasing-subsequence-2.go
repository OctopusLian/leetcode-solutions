func lengthOfLIS(nums []int) int {
    //二分查找
    if len(nums) == 0 {
        //考虑特殊情况
        return 0
    }
    heaps := []int{nums[0]}
    for i := 1; i < len(nums); i ++ {
        pos := binarySearch(heaps,0,len(heaps)-1,nums[i])
        if pos == -1 {
            heaps[0] = nums[i]
        }else if pos == len(heaps) {
            heaps = append(heaps,nums[i])  //找不到（pos == len(heaps)），那么在heaps中再添加一个数字
        }else {
            heaps[pos] = nums[i]  //找到了heaps中比当前数字大且序号最靠前的数字（pos=-1,pos有值），那么直接覆盖这个数字
        }
    }
    return len(heaps)
}

func binarySearch(nums []int,l,r int,target int) int {
    for {
        if l > r {
            break
        }
        mid := l + (r-l) >> 1
        if nums[mid] < target {
            l = mid+1
        }else {
            r = mid-1
        }
    }
    return l
}
