func merge(A []int, m int, B []int, n int)  {
    copy(A[m:], B[:])
    sort.Ints(A)
}
