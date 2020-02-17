func generateParenthesis(n int) []string {
    var ans []string
	var dfs func(s string, left, right int)
	
	dfs = func(s string, left, right int){
		if left > right { 	//从左向右构建，已使用的右括号的数量 必须小于 左括号
			return		//这条是关键规则
		}

		if left == 0 && right == 0 { // 左右括号都没有剩余了，构建完成
			ans = append(ans, s)
			return
		}

		if left > 0 {
			dfs(s + "(", left-1, right)
		}
		if right > 0 {
			dfs(s + ")", left, right-1)
		}
	}
	dfs("", n, n) //画一颗二叉树很容易理解
	return ans
}
