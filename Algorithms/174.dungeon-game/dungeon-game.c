int calculateMinimumHP(int** dungeon, int dungeonSize, int* dungeonColSize){
    //解法1，动态规划
    int n = dungeonSize, m = dungeonColSize[0];
    int dp[n + 1][m + 1];
    memset(dp, 0x3f, sizeof(dp));
    dp[n][m - 1] = dp[n - 1][m] = 1;
    for (int i = n - 1; i >= 0; --i) {
        for (int j = m - 1; j >= 0; --j) {
            int minn = fmin(dp[i + 1][j], dp[i][j + 1]);
            dp[i][j] = fmax(minn - dungeon[i][j], 1);
        }
    }
    return dp[0][0];
}