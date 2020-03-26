func maximum(a int, b int) int {
    c := a-b
    sa := sign(a) //如果a≥0，则返回1，否则返回0
    sb := sign(b) //如果b≥0，则返回1，否则返回0
    sc := sign(c)  //取决于a-b是否溢出

    use_sign_of_a := sa ^ sb
    use_sign_of_c := flip(sa ^ sb)
    k := use_sign_of_a * sa + use_sign_of_c * sc
    q := flip(k)  //q为k的反数

    return a * k + b * q
}

func flip(bit int) int {
    //将1翻转为0,,0翻转为1
    return 1^bit
}

func sign(a int) int {
    //如果是正数，就返回1；如果是负数，则返回0
    return flip((a >> 31) & 0x1)
}
