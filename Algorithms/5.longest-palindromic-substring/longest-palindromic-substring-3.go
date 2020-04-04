func longestPalindrome(s string) string {
    // 思路2，遍历字符串，找出以每个字符为中心的回文字符串有多长，选最长的返回
    length := len(s)
    getLen := func(i, j int) int {
        // 以s[i]s[j]为中心的最长回文字符串
        for i>=0 && j<length {
            if s[i] == s[j] {
                i--
                j++
            }else{
                return j - i - 1
            }
        }
        return j - i - 1
    }
    max := 0
    maxStart := 0
    for i:=0;i<length;i++ {
        if Max(getLen(i, i+1), getLen(i, i)) > max {
            max = Max(getLen(i, i+1), getLen(i, i))
            maxStart = i - (max-1)/2
        }
    }
    maxString := ""
    for i:=maxStart; i<maxStart+max; i++ {
        maxString += string(s[i])
    }
    return maxString
}

func Max(i, j int) int {
    if i >= j {
        return i 
    }else{
        return j
    }
}
