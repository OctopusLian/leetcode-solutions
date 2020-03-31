func sortArray(nums []int) []int {
    return insertSort(nums)
}

//插入排序
func insertSort(nums []int) []int {
    n := len(nums)
    for i := 1; i < n; i++ {
        e := nums[i]
        j := i
        for ; j > 0 && e < nums[j-1]; j-- {
            nums[j] = nums[j-1]
        }
        nums[j] = e
    }
    return nums
}
