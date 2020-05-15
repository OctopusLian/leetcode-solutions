class Solution {
    public int insertBits(int N, int M, int i, int j) {
        return ((N >> (j + 1)) << (j + 1)) | ((N >> i << i) ^ N) | (M << i);
    }
}

/**
 * 解答错误
 * 
 * 输入: 1143207437 1017033
 * 
 * 输出 2082995725 预期结果 2082885133
 * 
 * 
 */