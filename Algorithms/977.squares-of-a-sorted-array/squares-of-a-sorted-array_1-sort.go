func sortedSquares(A []int) []int {
	//解法1,直接排序
    ans := make([]int, len(A))
    for i, v := range A {
        ans[i] = v * v
    }
    sort.Ints(ans)
    return ans
}