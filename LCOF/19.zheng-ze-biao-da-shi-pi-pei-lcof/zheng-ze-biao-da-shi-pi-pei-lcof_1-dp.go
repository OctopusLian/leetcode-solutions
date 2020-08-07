var memory [][]bool

func isMatch(s string, p string) bool {
	//解法1，动态规划
	memory = make([][]bool, len(s)+1)
	for i := 0; i < len(memory); i++ {
		memory[i] = make([]bool, len(p)+1)
	}
	return dp(0, 0, s, p)
}

func dp(i int, j int, s string, p string) bool {
	if memory[i][j] {
		return memory[i][j]
	}
	if len(p) == 0 {
		return len(s) == 0
	}

	ok := len(s) != 0 && (s[0] == p[0] || p[0] == '.') //如果相等，或者为. 匹配成功
	var res bool
	if len(p) >= 2 && p[1] == '*' { //长度大于2，获取下一个元素且为*
		//1.去除x*（p[2:]）,继续递归匹配
		//2.去除字符(s[1:]),继续递归匹配
		res = dp(i, j+2, s, p[2:]) || (ok && dp(i+1, j, s[1:], p)) //只要一种情况成功，计算匹配成功
	} else {
		//普通匹配
		res = ok && dp(i+1, j+1, s[1:], p[1:])
	}
	memory[i][j] = res
	return res
}