func add(a int, b int) int {
    var sum,carry int
    for ;b != 0; {
	// 异或操作得无进位和
        sum = a ^ b
	// 与操作后移位得进位值
        carry = ((a & b) << 1)
	// 循环，直到进位为0
        a = sum
        b = carry
    }
    return a
}
