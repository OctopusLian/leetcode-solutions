class Solution {
public:
    bool canThreePartsEqualSum(vector<int>& A) {
        int sum = accumulate(A.begin(), A.end(), 0);  //累加求和
        if (sum % 3 != 0) {
            return false; //如果和取3的余不为0，返回false
        }
        int count = 0, subSum = 0;
        for (int i = 0; i < A.size(); i ++) {
            subSum += A[i];
            if (subSum == sum / 3) {
                count ++;
                subSum = 0;
            }
            if (count == 3) {
                return true;
            }
        }
        return false;
    }
};
