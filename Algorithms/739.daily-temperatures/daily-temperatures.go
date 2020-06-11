func dailyTemperatures(T []int) []int {
	//解法1，暴力
    length := len(T)
    ans := make([]int, length)
    next := make([]int, 101)
    for i := 0; i < 101; i++ {
        next[i] = math.MaxInt32
    }
    for i := length - 1; i >= 0; i-- {
        warmerIndex := math.MaxInt32
        for t := T[i] + 1; t <= 100; t++ {
            if next[t] < warmerIndex {
                warmerIndex = next[t]
            }
        }
        if warmerIndex < math.MaxInt32 {
            ans[i] = warmerIndex - i
        }
        next[T[i]] = i
    }
    return ans
}