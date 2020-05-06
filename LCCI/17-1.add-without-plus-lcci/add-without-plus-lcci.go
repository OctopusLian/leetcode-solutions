func add(a int, b int) int {
    if b == 0 {
        return a
    }
    sum := a ^ b  //两数相加，不进位
    carry := (a & b) << 1  //进位，但不对两数相加
    return add(sum,carry)  //以sum和carry为参数进行递归
}