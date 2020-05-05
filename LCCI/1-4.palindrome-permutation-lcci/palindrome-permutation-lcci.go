func canPermutePalindrome(s string) bool {
    // 只要判断是不是都是成双出现的字符就行
	data := make(map[string]int)
	count := 0
	for i := 0; i < len(s); i++ {
		data[string(s[i])]++
	}

	for _, value := range data {
		if value % 2 != 0 {
			count++
			if count >= 2 {
				return false
			}
		}
	}
	return true
}