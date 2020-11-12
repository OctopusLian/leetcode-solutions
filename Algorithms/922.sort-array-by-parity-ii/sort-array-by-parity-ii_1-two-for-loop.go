func sortArrayByParityII(A []int) []int {
	//解法1，两次遍历
    ans := make([]int, len(A))
    i := 0
    for _, v := range A {
        if v%2 == 0 {
            ans[i] = v
            i += 2
        }
    }
    i = 1
    for _, v := range A {
        if v%2 == 1 {
            ans[i] = v
            i += 2
        }
    }
    return ans
}