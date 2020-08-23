func rangeBitwiseAnd(m int, n int) int {
	//解法2，Brian-Kernighan 算法
    for m < n {
        n &= (n - 1)
    }
    return n
}