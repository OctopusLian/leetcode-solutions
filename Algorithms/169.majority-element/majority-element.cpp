class Solution {
public:
    int majorityElement(vector<int>& nums) {
        int res,cnt;
        res = 0;
        cnt = 0;
        for (int num : nums) {
            if (cnt == 0)
            {
                res = num;
                ++cnt;
            }
            else (num == res) ? ++cnt : --cnt;
        }
        return res;
    }
};
