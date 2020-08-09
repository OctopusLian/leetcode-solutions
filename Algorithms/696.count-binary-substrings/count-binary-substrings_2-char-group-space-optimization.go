func countBinarySubstrings(s string) int {
	//解法2，按字符分组的空间优化
	//用一个 last 变量来维护当前位置的前一个位置，这样可以省去一个 counts 数组的空间。
    //时间复杂度为 O(n)，空间复杂度为O(1)
    var ptr, last, ans int
    n := len(s)
    for ptr < n {
        c := s[ptr]
        count := 0
        for ptr < n && s[ptr] == c {
            ptr++
            count++
        }
        ans += min(count, last)
        last = count
    }

    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}