func moveZeroes(nums []int)  {
    index := 0
    //1，先将非0的数排到数组前面
    for _,v := range nums {
        if v != 0 {
            nums[index] = v
            index++
        }
    }
    //2，再排是0的数到nums数组后面
    for i:=index;i<len(nums);i++ {
        nums[i] = 0
    }
}
