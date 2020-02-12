func findRepeatNumber(nums []int) int {
    var nummap = make(map[int]bool)
    for _,num := range nums {
        if !nummap[num] {
            nummap[num] = true
        } else {
            return num
        }
    }

    return -1
}
