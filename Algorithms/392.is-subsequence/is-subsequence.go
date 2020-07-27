func isSubsequence(s string, t string) bool {
	//解法1，双指针法
    n, m := len(s), len(t)  
    i, j := 0, 0 //初始化两个指针i,j分别指向s和t初始化的位置
    for i < n && j < m {
        if s[i] == t[j] {
            i++  //匹配成功，则i,j同时右移
        }
        j++  //匹配失败，则j右移，i不变
    }
    return i == n  //最终如果i移动到s的末尾，就说明s是t的子序列
}