class Solution {
public:
    //1,数学+递归
    int lastRemaining(int n, int m) {
        return f(n,m);
    }

    int f(int n, int m) {
        if (n == 1)
            return 0;
        int x = f(n - 1, m);
        return (m + x) % n;
    }
};
