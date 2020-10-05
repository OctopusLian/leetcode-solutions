func numIdenticalPairs(nums []int) int {
    //暴力法
    count := 0
    for i,x := range nums {
        for j := i + 1;j < len(nums);j++ {
            if x == nums[j] {
                count++
            }
        }
    }
    return count
}