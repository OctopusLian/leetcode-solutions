func isPalindrome(x int) bool {
	xs := strconv.Itoa(x) // 整型转换为字符串
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		//i从左开始，j从右开始，i递增，j递减，逐个判断下标i和j对应的数字是否相等
		if xs[i] != xs[j] {
			return false
		}
	}
	return true
}
