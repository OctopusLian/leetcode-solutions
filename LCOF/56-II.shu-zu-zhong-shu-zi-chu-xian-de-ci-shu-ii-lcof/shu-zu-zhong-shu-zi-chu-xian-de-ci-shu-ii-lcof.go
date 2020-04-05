func singleNumber(nums []int) int {
    var numCountMap = make(map[int]int)
    for _,num := range nums {
        numCountMap[num]++
    }
    for num,count := range numCountMap {
        if count == 1 {
            return num
        }
    }
    return -1
}
