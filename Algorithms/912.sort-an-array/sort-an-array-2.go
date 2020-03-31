func sortArray(nums []int) []int {
    return selectSort(nums)
}

//选择排序
func selectSort(nums []int) []int {
    n := len(nums)
    l, r := 0, n-1
    for l < r {
        minIndex, maxIndex := l, r
        if nums[minIndex] > nums[maxIndex] {
            nums[minIndex], nums[maxIndex] = nums[maxIndex], nums[minIndex]
        }
        for i := l + 1; i < r; i++ {
            if nums[i] < nums[minIndex] {
                minIndex = i
            } else if nums[i] > nums[maxIndex] {
                maxIndex = i
            }
        }
        nums[l], nums[minIndex] = nums[minIndex], nums[l]
        nums[r], nums[maxIndex] = nums[maxIndex], nums[r]
        l++
        r--
    }
    return nums
}
