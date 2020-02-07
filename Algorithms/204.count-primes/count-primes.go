func countPrimes(n int) int {
    count := 0
    a := make([]bool, n)
    for i:=2;i<n;i++ {
        if a[i] {
            continue
        }
        for j:=i+i; j<n; j+=i {
            a[j] = true
        }
        count++
    }
    return count
}
