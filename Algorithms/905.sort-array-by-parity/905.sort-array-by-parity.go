func sortArrayByParity(A []int) []int {
	len := A.size - 1
	for i := 0;i < len;i++ {
		if A[i] % 2 == 1 {
			// A[i--] = temp
			// A[len--] = A[i--]
			A[i--],A[len--] = A[len--],A[i--]

		}
	}
	retuen A
}
