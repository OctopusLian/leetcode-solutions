func longestPalindrome(s string) int {
    dict := map[byte]bool{}
    res := 0
    for i := range s{
        if _,ok := dict[s[i]];ok{
            delete(dict,s[i]) //?
            res += 2 //?
        }else{
            dict[s[i]] = true
        }
    }
    if len(dict) > 0{
        res += 1
    }
    return res
}
