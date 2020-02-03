func reverse(x int) int {
    if x < 0 {
        //考虑x为负数的情况
        return -reverse(-x)
    }
    var re int
    for x != 0 {
        re = re*10 + x%10
        x = x/10
    }

    if re < 0x7fffffff {
        return re
    }
    return 0
}
