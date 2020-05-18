func validPalindrome(s string) bool {
	//解法1，双指针法
	return palindromeHelp(s,true)
}
func palindromeHelp(s string,isDelete bool)bool{
	left :=0
	right :=len(s)-1
	for left<right{
		if s[left]!=s[right]{  //当出现指针指向的字符不相同的时候
			if isDelete{
                //把左边的字符移动一位s[left+1:right+1]，或者把右边字符移动一位s[left:right]
				return palindromeHelp(s[left:right],false)||palindromeHelp(s[left+1:right+1],false)
			}else {
				return false
			}
		}
		left++
		right--
	}
	return true
}