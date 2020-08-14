func numberOf2sInRange(n int) int {
    if n < 2 {
        return 0
    }
    f := make([]int, 11)
    step := 1
    for i := 1; i < 11; i++ {
        f[i] = f[i-1]*10 + step
        step *= 10
    }
    var b []int
    tmp := n
    B := 1
    for tmp > 0 {
        b = append(b, tmp % 10)
        tmp /= 10
        if tmp > 0 {
            B *= 10
        }
    }
    var cnt int
    for i := len(b)-1; i >= 0; i-- {
        n %= B
        if b[i] > 0 {
            if b[i] < 2 {
                cnt += b[i]*f[i]
            } else if b[i] == 2 {
                cnt += b[i]*f[i]
                cnt += n+1
            } else {
                cnt += b[i]*f[i]
                cnt += B
            }
        }
        B /= 10
    }
    return cnt
}