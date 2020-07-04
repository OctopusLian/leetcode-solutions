func longestValidParentheses(s string) int {
	//解法1，栈
    maxAns := 0
    stack := []int{}
    stack = append(stack, -1)
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            //将它的下标放入栈中
            stack = append(stack, i)
        } else {
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {  //如果栈为空，说明当前的右括号为没有被匹配的右括号
                //将其下标放入栈中来更新「最后一个没有被匹配的右括号的下标」
                stack = append(stack, i)
            } else {  //如果栈不为空
                //当前右括号的下标减去栈顶元素即为「以该右括号为结尾的最长有效括号的长度」
                maxAns = max(maxAns, i - stack[len(stack)-1])
            }
        }
    }
    return maxAns
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}