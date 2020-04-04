class Solution {
public:
    int threeSumClosest(vector<int>& nums, int target) {
        if (nums.size() <= 3) {
            return accumulate(nums.begin(), nums.end(), 0);
        }
        sort (nums.begin(), nums.end());

        int result = 0, n = nums.size(), temp;
        result = nums[0] + nums[1] + nums[2];
        for (int i = 0; i < n - 2; ++i)
        {
            int j = i + 1, k = n - 1;
            while (j < k)
            {
                temp = nums[i] + nums[j] + nums[k];
                
                if (abs(target - result) > abs(target - temp))
                    result = temp;
                if (result == target)
                    return result;
                ( temp > target ) ? --k : ++j;
            }
        }
        return result;
    }
};
