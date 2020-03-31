func sortArray(nums []int) []int {
    return mergeSort(nums)
}

//归并排序
func mergeSort(nums []int) []int {
    mergeSortHelper(nums, 0, len(nums)-1)
    return nums
}

func merge(nums []int, l, m, r int) {
    auk := make([]int, r - l + 1)
    i, j := l, m+1
    k := 0
    for ; k < r - l + 1; k++ {
        if i > m {
            auk[k] = nums[j]
            j++
        } else if j > r {
            auk[k] = nums[i]
            i++
        } else if nums[i] < nums[j] {
            auk[k] = nums[i]
            i++
        } else {
            auk[k] = nums[j]
            j++
        }
    }
    copy(nums[l:r+1], auk)
}

func mergeSortHelper(nums []int, l, r int) {
    if r - l <= 15 {
        insertSortRange(nums, l, r)
        return
    }
    m := l + (r - l) / 2
    mergeSortHelper(nums, l, m)
    mergeSortHelper(nums, m+1, r)
    if nums[m] > nums[m+1] {
        merge(nums, l, m, r)
    }
}

//辅助其他排序的插入排序
func insertSortRange(nums []int, l, r int) {
    for i := l + 1; i <= r; i++ {
        e := nums[i]
        j := i
        for ; j > l && e < nums[j-1]; j-- {
            nums[j] = nums[j-1]
        }
        nums[j] = e
    }
}


