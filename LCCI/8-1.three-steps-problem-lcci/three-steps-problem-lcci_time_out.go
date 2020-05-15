func waysToStep(n int) int {
	//通俗解法
    if n < 0 {
        return 0
    } else if (n == 0) {
        return 1
    } else {
        return waysToStep(n-1) + waysToStep(n-2) + waysToStep(n-3)
    }
}

/*
执行结果：
超出时间限制
最后执行的输入：
61
*/