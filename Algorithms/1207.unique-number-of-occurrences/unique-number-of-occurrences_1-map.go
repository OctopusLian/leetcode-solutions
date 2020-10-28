func uniqueOccurrences(arr []int) bool {
	//解法1，哈希表
    cnts := map[int]int{}
    for _, v := range arr {
        cnts[v]++
    }
    times := map[int]struct{}{}
    for _, c := range cnts {
        times[c] = struct{}{}
    }
    return len(times) == len(cnts)
}