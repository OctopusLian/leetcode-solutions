class Solution {
public:
    vector<int> decompressRLElist(vector<int>& nums) {
        vector<int> res;
        for (int i = 0; 2 * i < nums.size(); i ++) 
            res.insert(res.end(), nums[2 * i], nums[2 * i + 1]);
        return res;
    }
};