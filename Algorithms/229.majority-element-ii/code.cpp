class Solution {
public:
    /**
     * @param nums: A list of integers
     * @return: The majority number occurs more than 1/3.
     */
    int majorityNumber(vector<int> nums) {
        if (nums.empty()) return -1;
        
        int k1 = 0, k2 = 0, c1 = 0, c2 = 0;
        for (auto n : nums) {
            if (!c1 || k1 == n) {
                k1 = n;
                c1++;
            } else if (!c2 || k2 == n) {
                k2 = n;
                c2++;
            } else {
                c1--;
                c2--;
            }
        }
        
        c1 = 0; 
        c2 = 0;
        for (auto n : nums) {
            if (n == k1) c1++;
            if (n == k2) c2++;
        }
        return c1 > c2 ? k1 : k2;
    }
};
