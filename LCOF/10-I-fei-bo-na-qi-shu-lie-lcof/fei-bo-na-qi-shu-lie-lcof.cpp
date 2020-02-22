class Solution {
public:
    int fib(int n) {
        //声明n+1+1大小的vector（第一个+1是要存储0至n共n+1个数据；第二个+1是考虑n==0的情况，保证v[1]不越界）
        vector<int> v(n + 1 + 1, 0); 
        v[1] = 1;
        for(int i = 2; i <= n; i++)
            v[i] = (v[i - 1] + v[i - 2]) % 1000000007;//注意别忘记取余
        return v[n];//直接返回最后一个数，即最终结果
    }
};
