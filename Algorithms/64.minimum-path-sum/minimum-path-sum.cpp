class Solution {
public:
    int minPathSum(vector<vector<int>>& grid) {
        //解法1，动态规划
        int m = grid.size(),n = grid[0].size();
        vector<vector<int>> dp(m+1,vector(n+1,99999)); dp[m][n-1] = 0;  // 边界条件

        for(int i=m-1;i>=0;--i){
            for(int j=n-1;j>=0;--j){
                dp[i][j] = min(dp[i+1][j],dp[i][j+1]) + grid[i][j]; // 转移方程
            }
        }
        return dp[0][0];
    }
};