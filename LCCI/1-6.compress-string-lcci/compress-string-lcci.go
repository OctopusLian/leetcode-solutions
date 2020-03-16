func compressString(S string) string {
    //考虑特殊情况
    if S == "" {
		return S
	}
	res := ""
	cur := S[0]
	count := 1
    // 顺序扫描，不用回溯，直接存储次数与重复字符，不等则存到结果中
	for i := 1; i < len(S); i++ {
		if cur != S[i] {
            //目前字符与上一个字符不重复，记录重复字符和出现次数，然后将这次出现的字符次数记为１
			res += string(cur)
			res += strconv.Itoa(count)
			cur = S[i]
			count = 1
		} else {
            //字符重复，累加出现次数
			count++
		}
	}
	res += string(cur)
	res += strconv.Itoa(count)
	if len(res) >= len(S) {
		return S
	}
	return res
}
