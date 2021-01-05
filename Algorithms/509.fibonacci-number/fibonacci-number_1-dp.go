func fib(n int) int {
	//解法1，动态规划
    if n < 2 {
        return n
    }
    p, q, r := 0, 0, 1
    for i := 2; i <= n; i++ {
        p = q
        q = r
        r = p + q
    }
    return r
}