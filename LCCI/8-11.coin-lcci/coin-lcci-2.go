func waysToChange(n int) int {
    ans := 0
    mod := 1000000007
    for i := 0; i * 25 <= n; i++ {
        rest := n - i * 25
        a, b := rest/10, rest%10/5
        ans = (ans + (a + 1) * (a + b + 1) % mod) % mod
    }
    return ans
}
