func getSum(a int, b int) int {
    //第一种方法，使用++和--
    if (a==0 || b==0){
        if a != 0 {
            return a
        } else {
            return b
        }
    }
    if (b < 0) {
        for ;b!=0; {
            b++
            a--
        }
        return a
    } else if (b > 0) {
        for ;b!=0; {
            a++
            b--
        }
        return a
    }

    return a
}
