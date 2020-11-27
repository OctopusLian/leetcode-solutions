func fourSumCount(A []int, B []int, C []int, D []int) int {
    var res int
    for _,i := range A {
        for _,j := range B {
            for _,k := range C {
                for _,l := range D {
                    if i + j + k + l == 0 {
                        res++
                    }
                }
            }
        }
    }

    return res
}

//有Bug，会超出时间限制