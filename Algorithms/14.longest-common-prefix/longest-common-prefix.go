func longestCommonPrefix(strs []string) string {
	//排除特殊情况
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	res := strs[0]               //获取字符串数组里的第一个元素
	for _, v := range strs[1:] { //从字符串数组第二个元素开始遍历
		var i int
		for ; i < len(v) && i < len(res); i++ { //遍历两数组里的元素
			if res[i] != v[i] { //做判断，如果不相等
				break //直接结束循环
			}
		}
		res = res[:i]
		if res == "" {
			return res //返回空
		}
	}

	return res
}
