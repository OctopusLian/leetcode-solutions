func trailingZeroes(n int) int {
    count := 0
    if n < 0 {
        return -1
    }
    for i := 5;n / i > 0;i *= 5 {
        count += n / i
    }
    return count
}