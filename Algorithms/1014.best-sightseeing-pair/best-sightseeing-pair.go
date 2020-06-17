func maxScoreSightseeingPair(A []int) int {
	//解法1，枚举法
    ans, mx := 0, A[0] + 0
    for j := 1; j < len(A); j++ {
        ans = max(ans, mx + A[j] - j)
        // 边遍历边维护
        mx = max(mx, A[j] + j)
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}