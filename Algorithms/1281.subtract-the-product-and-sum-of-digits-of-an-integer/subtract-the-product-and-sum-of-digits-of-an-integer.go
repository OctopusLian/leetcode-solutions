func subtractProductAndSum(n int) int {
    add := 0
    mul := 1
    for ;n > 0; {
        digit := n % 10
        n /= 10
        add += digit
        mul *= digit
    }
    return mul - add;
}