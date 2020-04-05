func singleNumbers(nums []int) []int {
    var result []int
    var numCountMap = make(map[int]int)
    for _,num := range nums {
        numCountMap[num]++
    }
    for num,count := range numCountMap {
        if count == 1 {
            result = append(result,num)
        }
    }
    return result
}
