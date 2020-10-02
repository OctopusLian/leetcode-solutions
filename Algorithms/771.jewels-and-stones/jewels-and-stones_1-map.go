func numJewelsInStones(J string, S string) int {
    //解法1，map
    jewels := make(map[rune]bool)
    for _, j := range J {
        jewels[j] = true
    }
    count := 0
    for _, s := range S {
        if jewels[s] {
            count ++
        }
    }
    return count
}
