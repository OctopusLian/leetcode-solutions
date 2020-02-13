func rotate(nums []int, k int)  {
    //使用额外的数组
    a := make([]int,len(nums))
    for i:=0;i<len(nums);i++ {
        a[(i+k) % len(nums)] = nums[i]
    }
    for i:=0;i<len(nums);i++ {
        nums[i] = a[i]
    }
    
}
