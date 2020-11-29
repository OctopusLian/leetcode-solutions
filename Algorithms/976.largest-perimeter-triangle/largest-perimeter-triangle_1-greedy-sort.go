func largestPerimeter(A []int) int {
	//解法1，贪心+排序
    sort.Ints(A)
    for i := len(A) - 1; i >= 2; i-- {
        if A[i-2]+A[i-1] > A[i] {
            return A[i-2] + A[i-1] + A[i]
        }
    }
    return 0
}