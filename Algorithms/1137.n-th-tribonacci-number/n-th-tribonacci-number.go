func tribonacci(n int) int {
    if n < 3 {
        if n == 0 {
            return 0
        } else {
            return 1
        }
    }

    tmp,x,y,z := 0,0,1,1
    for i := 3;i <= n;i++ {
        tmp = x + y + z
        x = y
        y = z
        z = tmp
    }

    return z
}