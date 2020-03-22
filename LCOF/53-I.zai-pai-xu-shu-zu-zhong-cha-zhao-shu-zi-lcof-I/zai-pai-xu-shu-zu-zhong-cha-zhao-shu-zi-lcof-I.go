func search(nums []int, target int) int {
    TargetCount := make(map[int]int)
    for _,num := range nums {
        TargetCount[num]++
    }

    return TargetCount[target]
}
