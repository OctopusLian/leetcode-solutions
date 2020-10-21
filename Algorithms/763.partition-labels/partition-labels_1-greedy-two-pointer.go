func partitionLabels(S string) []int {
    //解法1，贪心＋双指针
    var partition []int
    lastPos := [26]int{}
    for i, c := range S {
        lastPos[c-'a'] = i
    }
    start, end := 0, 0
    for i, c := range S {
        if lastPos[c-'a'] > end {
            end = lastPos[c-'a']
        }
        if i == end {
            partition = append(partition, end-start+1)
            start = end + 1
        }
    }
    return partition
}