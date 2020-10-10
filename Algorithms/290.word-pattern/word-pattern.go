func wordPattern(pattern string, str string) bool {
	// 对于字符串使用split为切片
	// 对于pattern取每个字符串
	// 然后依次判断每个字符串在后面的字符串的位置如果不同则为false
	sp := strings.Split(str, " ")
	if len(sp) == 1 {
		// 如果sp为一个 这时候判断pattern的个数如果为1
		// 那么就直接返回true 否则为false
		return len(pattern) == 1
	}
	for i := 0; i < len(pattern); i++ {
		if strings.Index(pattern[i+1:], string(pattern[i])) != find(sp[i+1:], sp[i]) {
			return false
		}
	}
	return true
}
// 寻找索引的函数
func find(lst []string, v string) (index int) {
	for i, item := range lst {
		if v == item {
			index = i
			return
		}
	}
	// 与strings.Index一致没有找到时需要返回-1
	return -1
}