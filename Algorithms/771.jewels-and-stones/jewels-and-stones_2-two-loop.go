func numJewelsInStones(J string, S string) int {
	//解法二，暴力求解，两个for循环
    count := 0
    for _,s := range S {
        for _,j := range J {
            if s == j {
                count++
                break
            }
            
        }
    }

    return count
}