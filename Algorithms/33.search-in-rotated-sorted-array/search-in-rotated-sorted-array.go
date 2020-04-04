func search(nums []int, target int) int {
    for index,num := range nums {
        if num == target {
            return index
        }
    }
    return -1
}
