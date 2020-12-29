func splitIntoFibonacci(S string) []int {
    /*
        栈保存每一步状态
    */
    stack  := make([]int, len(S))
    top := -1

    var search func(string) bool
    search = func(s string) bool{
        // 模板第一步

        if len(s) == 0 {
            /*要么全是0， 要么最后一个不是零*/
            if top >= 2 && (stack[top] !=0 || stack[top] + stack[top-1] == 0){
                return true
            } else {
                return false
            }
        }

        // 第二步 1
        /*  当前可行状态有三种可能
                1. 如果是0， 则只有一种状态
                2. 非零且栈中大于等于两个元素，则当前元素(状态)需要等于上两个元素之和，也只有一种状态
                3. 非零且栈中元素小于两个元素，则当前元素(状态)可任取,需要遍历
            需要分别处理
        */
 
        /* 如果当前为零 则必取,状态数为1*/ 
        if s[0] == '0' {
            // 第三步
            top ++
            stack[top] = 0
            // 第四步
            if search(s[1:]) {
                return true
            } else {
                /*回溯*/
                top--
                return false
            }
        }

        // 第二步 2
        if top >= 1 {
            for i:=1; i<=len(s); i++ {
                v,_ := strconv.Atoi(s[0:i])
                /* 剪枝 */
                if v > ((1<<31)-1) || v > stack[top] + stack[top-1] {
                    break
                }

                /* 找到了唯一的元素(状态) */
                if v == stack[top] + stack[top-1] {
                    // 第三步
                    top++
                    stack[top] = v

                    // 第四步
                    if search(s[i:]) {
                        return true
                    }
                    /* 回溯 */
                    top--
                    break
                }
            }
        // 第二步 3
        } else {
            for i:=1;i<=len(s);i++ {
                v,_ := strconv.Atoi(s[0:i])
                /* 剪枝 */
                if v > ((1 << 31)-1) {
                    break
                }
                // 第三步
                top ++
                stack[top] = v
                // 第四步
                if search(s[i:]) {
                    return true
                }
                /* 遍历下一个状态 */
                top--
            }
        }

        // 第五步
        return false
    }

    if search(S) {
        return stack[:top+1]
    } else {
        return []int{}
    }
}