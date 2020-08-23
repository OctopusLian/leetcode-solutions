func rangeBitwiseAnd(m int, n int) int {
	//解法1，位移
    shift := 0
    for m < n {
        m, n = m >> 1, n >> 1
        shift++
    }
    return m << shift
}