func fib(n int) int {
    a := make([]int,n+2)
    a[1] = 1
    for i := 2;i <= n;i++ {
        a[i] = (a[i-1] + a[i-2]) % 1000000007
    }
    return a[n]
}
