func removeKdigits(num string, k int) string {
	//解法1，贪心+单调栈
    stack := []byte{}
    for i := range num {
        digit := num[i]
        for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
            k--
        }
        stack = append(stack, digit)
    }
    stack = stack[:len(stack)-k]
    ans := strings.TrimLeft(string(stack), "0")
    if ans == "" {
        ans = "0"
    }
    return ans
}