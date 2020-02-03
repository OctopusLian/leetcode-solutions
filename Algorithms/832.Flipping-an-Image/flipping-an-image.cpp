class Solution {
public:
    vector<vector<int>> flipAndInvertImage(vector<vector<int>>& A) {
        size_t len = A.size();  //获得二进制数组的长度
        for (int i = 0; i < A.size(); i++)
        {
            reverse(A[i].begin(),A[i].end());  //执行翻转（逆序）二进制矩阵的操作
            for (int j = 0; j < A[i].size(); j++)
            {
                A[i][j] = 1 - A[i][j];  //执行反转二进制矩阵的操作
            }
        }
        return A;        
    }
};
