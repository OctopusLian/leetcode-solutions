func numberOfSteps (num int) int {
    var ans int
    for ;; ans++ {
        if num == 0 {
            break
        } else if num & 1 == 1 {
            num--
        } else {
            num >>= 1
        }
    }
    return ans
}