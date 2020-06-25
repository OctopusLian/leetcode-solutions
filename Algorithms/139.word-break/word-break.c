unsigned long long Hash(char* s, int l, int r) {
    unsigned long long value = 0;
    for (int i = l; i < r; i++) {
        value = value * 2333ull;
        value += s[i] - 'a' + 1;
    }
    return value;
}
bool query(unsigned long long* rec, int len_rec, unsigned long long x) {
    for (int i = 0; i < len_rec; i++) {
        if (rec[i] == x) return true;
    }
    return false;
}
bool wordBreak(char* s, char** wordDict, int wordDictSize) {
    //解法1，动态规划
    unsigned long long rec[wordDictSize + 1];
    for (int i = 0; i < wordDictSize; i++) {
        rec[i] = Hash(wordDict[i], 0, strlen(wordDict[i]));
    }
    int len_s = strlen(s);
    bool dp[len_s + 1];
    memset(dp, 0, sizeof(dp));
    dp[0] = true;
    for (int i = 1; i <= len_s; i++) {
        for (int j = 0; j < i; j++) {
            if (dp[j] && query(rec, wordDictSize, Hash(s, j, i))) {
                dp[i] = true;
                break;
            }
        }
    }
    return dp[len_s];
}