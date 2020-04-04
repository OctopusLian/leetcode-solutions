func longestPalindrome(s string) string {
    res := "" //存放结果
    temp := ""  //存放子串
    for i:=0;i<len(s);i++ {
        for j:=i;j<len(s);j++ {
            byte_temp := []byte(temp)
            byte_temp = append(byte_temp, []byte(s)[j])
            temp = string(byte_temp)
            tem := temp //存放子串反转的结果
            //反转
            tem = reverse(temp)
            if tem == temp {
                if len(res)>len(temp) {
                    res = res
                } else {
                    res = temp
                }
            }
        }
        temp = ""
    }

    return res
}

func reverse(str string) string {
    rs := []rune(str)
    len := len(rs)
    var tt []rune

    tt = make([]rune, 0)
    for i := 0; i < len; i++ {
        tt = append(tt, rs[len-i-1])
    }
    return string(tt[0:])
}

// 有Bug
