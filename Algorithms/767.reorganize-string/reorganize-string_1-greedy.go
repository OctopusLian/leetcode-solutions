func reorganizeString(S string) string {
	//解法1，基于计数的贪心算法
    n := len(S)
    if n <= 1 {
        return S
    }

    cnt := [26]int{}
    maxCnt := 0
    for _, ch := range S {
        ch -= 'a'
        cnt[ch]++
        if cnt[ch] > maxCnt {
            maxCnt = cnt[ch]
        }
    }
    if maxCnt > (n+1)/2 {
        return ""
    }

    ans := make([]byte, n)
    evenIdx, oddIdx, halfLen := 0, 1, n/2
    for i, c := range cnt[:] {
        ch := byte('a' + i)
        for c > 0 && c <= halfLen && oddIdx < n {
            ans[oddIdx] = ch
            c--
            oddIdx += 2
        }
        for c > 0 {
            ans[evenIdx] = ch
            c--
            evenIdx += 2
        }
    }
    return string(ans)
}