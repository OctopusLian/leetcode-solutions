class Solution {
public:
    int mySqrt(int x) {
        if (x <= 0) return 0;

    int lb = 0, ub = x;
    while (lb + 1 < ub) {
        long long mid = lb + (ub - lb) / 2; 
        if (mid * mid == x) return mid; 
        if (mid * mid < x) lb = mid;
        else ub = mid;
    }
    return lb;
    }
};

//有Bug
