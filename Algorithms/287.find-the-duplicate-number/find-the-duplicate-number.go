func findDuplicate(nums []int) int {
    var numsmap = make(map[int]bool)
    for _,num := range nums {
        if !numsmap[num] {
            //只出现一次
            numsmap[num] = true
        } else {
            return num
        }
    }
    return -1
}