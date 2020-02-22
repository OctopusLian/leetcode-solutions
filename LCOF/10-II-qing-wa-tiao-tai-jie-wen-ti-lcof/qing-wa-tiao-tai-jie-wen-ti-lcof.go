func numWays(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    a := make([]int,n+1)
    a[0] = 1
    a[1] = 1
    for i := 2;i <= n;i++ {
        a[i] = (a[i-1] + a[i-2]) % 1000000007
    }
    return a[n]
}
