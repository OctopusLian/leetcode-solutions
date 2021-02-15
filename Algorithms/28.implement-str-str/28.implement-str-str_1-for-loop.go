func strStr(haystack string, needle string) int {
	//考虑特殊情况
	if len(haystack) == 0 && len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	len_h := len(haystack) //获取haystack字符串的长度
	len_n := len(needle)   //获取needle字符串的长度
	for i := 0; i < len_h-len_n+1; i++ {
		j := 0 //子串每次都要重头开始遍历
		for ; j < len_n; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len_n {
			return i
		}

	}

	return -1
}
