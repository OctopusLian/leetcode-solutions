func gcdOfStrings(str1 string, str2 string) string {
    l1,l2 := len(str1),len(str2)
	if l1 < l2 {
		return gcdOfStrings(str2, str1)
	}

	if l1 == l2 {
		if str1 == str2 {
            //如果两个字符串完全相等，则返回该字符串
			return str1
		}
        //否则返回空
		return ""
	}
    //如果两个字符串长度不相等，则让str1长度大于str2
	for i := 0; i < l2; i++ {
		if str1[i] != str2[i] {
			return ""
		}
	}

	return gcdOfStrings(str1[l2:], str2)
}
