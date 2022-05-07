/**
 * @param n: An integer
 * @return: An integer
 */
 func climbStairs (n int) int {
    // write your code here
    if n <= 2 {
        //小于2，直接返回值
        return n
    }

    FirstNum := 1
    SecondNum := 2
    result := 0
    for i:=3;i<=n;i++ {
        result = FirstNum + SecondNum  //当前层 = 倒数第一层 + 倒数第二层
        FirstNum = SecondNum
        SecondNum = result
    }

    return result
}