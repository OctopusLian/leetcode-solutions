func sortArray(nums []int) []int {
    return quickSort(nums)
}

//三路快排
func quickSort(nums []int) []int {
    rand.Seed(time.Now().Unix())
    return quickSortHelper(nums, 0, len(nums)-1)
}

func quickSortHelper(nums []int, l, r int) []int {
    if r - l <= 15 {
        insertSortRange(nums, l, r)
        return nums
    }
    randIndex := rand.Intn(r-l+1) + l
    nums[l], nums[randIndex] = nums[randIndex], nums[l]
    lt, rt, i := l, r+1, l+1
    v := nums[l]
    for i < rt {
        if nums[i] < v {
            lt++
            nums[lt], nums[i] = nums[i], nums[lt]
            i++
        } else if nums[i] > v {
            rt--
            nums[rt], nums[i] = nums[i], nums[rt]
        } else {
            i++
        }
    }
    nums[l], nums[lt] = nums[lt], nums[l]
    quickSortHelper(nums, l, lt-1)
    quickSortHelper(nums, rt, r)
    return nums
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
