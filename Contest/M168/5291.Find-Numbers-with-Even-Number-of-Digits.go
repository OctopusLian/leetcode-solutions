func findNumbers(nums []int) int {
    count := 0
    for _,value := range nums {
        if len(strconv.Itoa(value)) % 2 == 0 {
            count++
        }
    }
    
    return count
}
