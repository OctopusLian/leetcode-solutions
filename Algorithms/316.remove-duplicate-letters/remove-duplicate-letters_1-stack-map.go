func removeDuplicateLetters(s string) string {
    //stack + map
    n := len(s)
    indexMap, seenMap := map[byte]int{}, map[byte]bool{} // 一个存储最后出现的下标，一个是存储是否已经出现过了
    stack := []byte {}
    for i := 0; i < n; i++ {
        indexMap[s[i]] = i
    }
    for i := 0; i < n; i++ {
        tempByte := s[i]
        if _, ok := seenMap[tempByte]; !ok { // 没有被加入过结果栈中
            for len(stack) > 0 && stack[len(stack) - 1] > tempByte && indexMap[stack[len(stack) - 1]] > i  { // 后面还会出现，从当前结果栈弹出，要后面的保证字典序最小, 并且结果栈也没见过这个字符
                size := len(stack) - 1
                delete(seenMap, stack[size])
                stack = stack[:size]
            } 
            stack = append(stack, tempByte)
            seenMap[tempByte] = true
        }
    }
    res :=  ""
    for _, val := range(stack) {
        res += string(val)
    }
    return res
}