func longestCommonPrefix(strs []string) string {
    //排除特殊情况
    if len(strs) == 0 {
        return ""
    }
    if len(strs) == 1{
        return strs[0]
    }
    res := strs[0]  //获取数组里的第一个元素
    for _,v := range strs[1:] {
        var i int
        for ;i < len(v) && i < len(res);i++ {
            if res[i] != v[i] {
                break
            }
        }
        res = res[:i]
        if res == "" {
            return res
        }
    }

    return res
}
