func fib(n int) int {
	//解法2，递归
    if n <= 1 {
        return n
    }
    return fib(n-1) + fib(n-2)
}