func firstUniqChar(s string) byte {
    res := make([]int, 26)  //建立一个长度为26的整数型数组
    //按顺序统计每个字符出现的次数
    for _, si := range s {
		res[si - 'a']++
	}
    //按顺序查找
	for _, si := range s {
		if res[si - 'a'] == 1 {
			return byte(si)
		}
	}
	return ' '
}
