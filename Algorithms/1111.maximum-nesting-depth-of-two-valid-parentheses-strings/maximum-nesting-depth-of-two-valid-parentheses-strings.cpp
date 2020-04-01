class Solution {
public:
    vector<int> maxDepthAfterSplit(string seq) {
        //1，use stack
        int d = 0;
        vector<int> ans;
        for (char& c : seq)
            if (c == '(') {
                ++d;  //增加深度
                ans.push_back(d % 2);
            }
            else {
                ans.push_back(d % 2);
                --d;  //减少深度
            }
        return ans;
    }
};
