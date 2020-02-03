func isValid(s string) bool {
    //考虑空字符串的特殊情况
    if s == "" {
		return true
	}
    //定义一个栈
	stack := make([]int32, len(s))
    length := 0
    //判断括号是否匹配
    for _,v := range s {
        if v == '(' || v == '[' || v == '{' {
            //左括号，入栈
            stack[length] = v
            length++
        } else {
            //右括号，比较栈顶，匹配则移除，不匹配就返回false
            if length == 0 {
                return false
            }
            if (v == ')' && stack[length-1] == '(') || (v == ']' && stack[length-1] == '[') || (v == '}' && stack[length-1] == '{') {
                length--
            } else {
                return false
            }
        }
    }

    return length == 0
}
