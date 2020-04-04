func longestPalindrome(s string) string {
    // 思路1：找出每一个子字符串，判断其是不是回文字符串
    length := len(s)
    if length == 0 {
        return ""
    }
    max := 0
    maxl := 0
    maxr := 0
    for i:=0;i<length;i++ {
        for j:=i+1; j<length; j++ {
            ret := isPalindromic(s, i, j)
            if ret {
                if max < (j-i+1) {
                    max = j - i + 1
                    maxl = i
                    maxr = j
                }
            }
        }
    }
    maxString := ""
    for i:=maxl; i<=maxr; i++ {
        maxString += string(s[i])
    }
    return maxString
}

func isPalindromic(s string, l, r int) bool {
    for l < r {
        if s[l] == s[r] {
            l++
            r--
        }else{
            return false
        }
       
    }
    return true
}
