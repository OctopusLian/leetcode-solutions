func heightChecker(heights []int) int {
	//解法1，桶排序
    count := make([]int, 101)
    for _, h := range heights {
        count[h]++
    }

    ret := 0
    j := 0
    for i := 1; i < len(count); i++ {
        for ; count[i] > 0; count[i]-- {
            if heights[j] != i {
                ret++
            }
            j++
        }
    }
    return ret
}