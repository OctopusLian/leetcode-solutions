func kidsWithCandies(candies []int, extraCandies int) []bool {
	//解法1，枚举
    n := len(candies)
    maxCandies := 0
    for i := 0; i < n; i++ {
        maxCandies = max(maxCandies, candies[i])
    }
    ret := make([]bool, n)
    for i := 0; i < n; i++ {
        ret[i] = candies[i] + extraCandies >= maxCandies
    }
    return ret
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}