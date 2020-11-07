func removeOuterParentheses(S string) string {
	//解法1，栈
    stack := make([]rune, 0)
    level := 0
    for _, v := range S {
        if v == '(' {
            level++
        } 
        if level > 1 {
            stack = append(stack, v)
        }
        if v == ')' {
            level--
        }
	}
	
    return string(stack)
}