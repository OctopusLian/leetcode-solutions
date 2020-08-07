func getValidT9Words(num string, words []string) []string {
	//解法1，map
    strMap := make(map[byte]byte,10)
    strMap['a'] = '2'
    strMap['b'] = '2'
    strMap['c'] = '2'
    strMap['d'] = '3'
    strMap['e'] = '3'
    strMap['f'] = '3'
    strMap['g'] = '4'
    strMap['h'] = '4'
    strMap['i'] = '4'
    strMap['j'] = '5'
    strMap['k'] = '5'
    strMap['l'] = '5'
    strMap['m'] = '6'
    strMap['n'] = '6'
    strMap['o'] = '6'
    strMap['p'] = '7'
    strMap['q'] = '7'
    strMap['r'] = '7'
    strMap['s'] = '7'
    strMap['t'] = '8'
    strMap['u'] = '8'
    strMap['v'] = '8'
    strMap['w'] = '9'
    strMap['x'] = '9'
    strMap['y'] = '9'
    strMap['z'] = '9'
    res := make([]string,0,10)
    lenth := len(num)
    lenthA := len(words)
    for i:=0;i<lenthA;i++{
        if len(words[i]) != lenth{
            continue
        }
        for j:=0;j<lenth;j++{
            if strMap[words[i][j]] != num[j]{
                break
            }
            if j == lenth-1{
                res = append(res,words[i])
            }
        }
    }
    return res
}