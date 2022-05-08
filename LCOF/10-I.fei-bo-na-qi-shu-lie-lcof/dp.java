/*
 * @Description: 
 * @Author: neozhang
 * @Date: 2022-05-08 16:31:01
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-08 16:31:04
 */
class Solution {
    public int fib(int n) {
        //边界处理
        if (n == 0) return 0;

        //初始化数组dp
        int[] dp = new int[n + 1];

        //由于F(0) = 0，所以 dp[0] = 0
        dp[0] = 0;

        //由于F(1) = 1，所以 dp[1] = 1
        dp[1] = 1;

        //通过for循环来填充dp
        for (int i = 2;i <= n;i++) {
            //dp[i]的计算规则
            //F(N) = F(N - 1) + F(N - 2)
            dp[i] = dp[i-1] + dp[i-2];
            //答案需要取模 1e9+7 (1000000007)
            dp[i] %= 1000000007;
        }

        //返回结果
        return dp[n];
    }
}