func minPatches(nums []int, n int) int {
    var patches int
    for i, x := 0, 1; x <= n; {
        if i < len(nums) && nums[i] <= x {
            x += nums[i]
            i++
        } else {
            x *= 2
            patches++
        }
    }
    return patches
}