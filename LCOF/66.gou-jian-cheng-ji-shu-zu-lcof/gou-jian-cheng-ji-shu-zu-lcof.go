func constructArr(a []int) []int {
    n := len(a)
    if n == 0 {
        return []int{}
    }
    b := make([]int, n)
    //先左乘，把当前左边的数相乘
    left := 1
    for i := 0; i < n; i++ {
        b[i] = left
        left *= a[i]
    }
    //再右乘，把当前右边的数相乘
    right := a[n-1]
    for i := n - 2; i >= 0; i-- {
        b[i] *= right
        right *= a[i]
    }
    return b
}
