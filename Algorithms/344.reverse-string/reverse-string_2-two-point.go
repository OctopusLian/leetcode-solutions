func reverseString(s []byte) {
	//解法２，双指针法
    for left, right := 0, len(s)-1; left < right; left++ {
        s[left], s[right] = s[right], s[left]
        right--
    }
}