func minCount(coins []int) int {
    count := 0
    for _,coin := range coins {
        if (coin % 2 == 0) {
            count = count + coin/2
        } else {
            count = 1 + count + coin/2
        }
    }
    return count
}
